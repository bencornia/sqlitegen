package codegen

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/tools/imports"
)

type schema struct {
	Name    string
	Columns []*column
}

type column struct {
	Name         string
	Type         string
	NotNull      bool
	IsPrimaryKey bool
}

func catch(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type closable interface {
	Close() error
}

func catchClosable(c closable) {
	catch(c.Close())
}

func isValidSchema(db *sql.DB, tableName string) (bool, error) {
	query := fmt.Sprintf(`
		with flags as (
			select
				(
					select strict
					from pragma_table_list('%s')
				) as is_strict_table,
				(
					select sum("notnull" = 0) = 0
					from pragma_table_info('%s')
				) as has_not_null_columns,
				(
					select count(*) = 1
					from pragma_table_info('%s')
					where name = 'id'
						and type = 'INTEGER'
						and pk = 1
				) as has_valid_pk,
				(
					select count(*) = 1
					from pragma_table_info('%s')
					where name = 'created_at'
						and type = 'TEXT'						
				) as has_valid_created_at,
				(
					select count(*) = 1
					from pragma_table_info('%s')
					where name = 'updated_at'
						and type = 'TEXT'
				) as has_valid_updated_at
		)
		select is_strict_table
			and has_not_null_columns
			and has_valid_pk
			and has_valid_created_at
			and has_valid_updated_at
		from flags;
	`, tableName, tableName, tableName, tableName, tableName)

	var isValid bool
	err := db.QueryRow(query).Scan(&isValid)
	return isValid, err
}

func getSchemas(db *sql.DB) ([]*schema, error) {
	tableNames, err := db.Query("select name from sqlite_master where type = 'table'")
	if err != nil {
		return nil, err
	}

	defer catchClosable(tableNames)

	var schemas []*schema
	for tableNames.Next() {
		var tableName string
		err = tableNames.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		ok, err := isValidSchema(db, tableName)
		if err != nil {
			return nil, err
		}

		if !ok {
			continue
		}

		query := fmt.Sprintf("select name, type, `notnull`, pk from pragma_table_info('%s')", tableName)
		columns, err := db.Query(query)
		catch(err)

		s := &schema{Name: tableName, Columns: []*column{}}
		for columns.Next() {
			var col column
			err = columns.Scan(&col.Name, &col.Type, &col.NotNull, &col.IsPrimaryKey)
			if err != nil {
				return nil, err
			}

			s.Columns = append(s.Columns, &col)
		}

		catchClosable(columns)
		schemas = append(schemas, s)
	}

	return schemas, nil
}

func Generate(dsn string, packageName string, writer io.Writer) {
	// Step 1: Get database
	db, err := sql.Open("sqlite3", dsn)
	catch(err)

	// Step 2: Check for existing tables
	var exists bool
	err = db.QueryRow("select count(*) > 0 from sqlite_master").Scan(&exists)
	catch(err)
	if !exists {
		catch(fmt.Errorf("no tables in %s", dsn))
	}

	// Step 3:
	schemas, err := getSchemas(db)
	catch(err)
	if len(schemas) == 0 {
		catch(fmt.Errorf("no supported schemas"))
	}

	// Step 4: Register template functions
	funcs := template.FuncMap{
		"getTag":      getTag,
		"pascalCase":  pascalCase,
		"camelCase":   camelCase,
		"getType":     getType,
		"columnNames": columnNames,
		"join":        joinItems,
		"map":         mapItems,
		"filter":      filterItems,
		"backtick":    backtick,
	}

	// Step 5: Execute template
	var data struct {
		PackageName string
		Schemas     []*schema
	}

	data.PackageName = packageName
	data.Schemas = schemas

	var buf bytes.Buffer
	tmpl := template.Must(template.New("").Funcs(funcs).ParseFiles("internal/codegen/template.tmpl"))
	err = tmpl.ExecuteTemplate(&buf, "base", data)
	catch(err)

	// Step 6: Format code
	opts := &imports.Options{
		Fragment:   false,
		AllErrors:  false,
		Comments:   true,
		TabIndent:  false,
		FormatOnly: false,
	}

	formatted, err := imports.Process("foo.go", buf.Bytes(), opts)
	catch(err)

	// Step 7: Write file
	_, err = writer.Write(formatted)
	catch(err)
}
