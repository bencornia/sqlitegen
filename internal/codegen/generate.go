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

func Generate(dsn string, packageName string, writer io.Writer) {
	// Step 1) Get database
	db, err := sql.Open("sqlite3", dsn)
	catch(err)

	// Step 2) Check for existing tables
	var exists bool
	err = db.QueryRow("select count(*) > 0 from sqlite_master").Scan(&exists)
	catch(err)

	if !exists {
		catch(fmt.Errorf("no tables in %s", dsn))
	}

	// Step 3) Get schemas
	tableNames, err := db.Query("select name from sqlite_master where type = 'table'")
	catch(err)

	var data struct {
		PackageName string
		Schemas     []*schema
	}

	var schemas []*schema
	defer func(rows *sql.Rows) {
		catch(rows.Close())
	}(tableNames)

	for tableNames.Next() {
		var tableName string
		err = tableNames.Scan(&tableName)
		catch(err)

		query := fmt.Sprintf("select name, type, `notnull`, pk from pragma_table_info('%s')", tableName)
		columns, err := db.Query(query)
		catch(err)

		defer func(rows *sql.Rows) {
			catch(rows.Close())
		}(columns)

		s := &schema{Name: tableName, Columns: []*column{}}
		for columns.Next() {
			var col column
			err = columns.Scan(&col.Name, &col.Type, &col.NotNull, &col.IsPrimaryKey)
			catch(err)

			s.Columns = append(s.Columns, &col)
		}

		// Ensure that the columns include id, updated_at, created_at
		missingPrimaryKey := true
		missingUpdatedAt := true
		missingCreatedAt := true
		for _, col := range s.Columns {
			switch col.Name {
			case "id":
				missingPrimaryKey = !col.IsPrimaryKey
			case "updated_at":
				missingUpdatedAt = false
			case "created_at":
				missingCreatedAt = false
			}
		}

		isMissingFields := missingPrimaryKey || missingUpdatedAt || missingCreatedAt
		if isMissingFields {
			continue
		}

		schemas = append(schemas, s)
	}

	if len(schemas) == 0 {
		catch(fmt.Errorf("no supported schemas"))
	}

	data.PackageName = packageName
	data.Schemas = schemas

	// Step 4) Register template functions
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

	// Step 5) Execute template
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

	// Write file
	_, err = writer.Write(formatted)
	catch(err)
}
