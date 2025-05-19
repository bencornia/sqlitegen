package codegen

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"
	"text/template"
	"unicode"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/tools/imports"
)

type schema struct {
	Name    string
	Columns []*column
}

type column struct {
	Name    string
	Type    string
	NotNull bool
}

func catch(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func Generate(dsn string, writer io.Writer) {
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

	var schemas []*schema
	defer func(rows *sql.Rows) {
		catch(rows.Close())
	}(tableNames)

	for tableNames.Next() {
		var tableName string
		err = tableNames.Scan(&tableName)
		catch(err)

		query := fmt.Sprintf("select name, type, `notnull` from pragma_table_info('%s')", tableName)
		columns, err := db.Query(query)
		catch(err)

		defer func(rows *sql.Rows) {
			catch(rows.Close())
		}(columns)

		s := &schema{Name: tableName, Columns: []*column{}}
		for columns.Next() {
			var col column
			err = columns.Scan(&col.Name, &col.Type, &col.NotNull)
			catch(err)

			s.Columns = append(s.Columns, &col)
		}

		schemas = append(schemas, s)
	}

	// Step 4) Register template functions
	funcs := template.FuncMap{
		"jsonTag": func(col *column) string {
			return fmt.Sprintf("`json:\"%s\"`", col.Name)
		},
		// PascalCase
		"pascalCase": func(val string) string {
			var (
				result = ""
				runes  = []rune(val)
				i      = 0
				upper  = false
			)

			for i < len(runes) {
				if runes[i] == '_' {
					i++
					upper = true
					continue
				}

				char := runes[i]
				if i == 0 || (runes[i-1] == '_' && upper) {
					char = unicode.ToUpper(char)
				}

				result += string(char)
				i++
				upper = false
			}

			return result
		},
		// camelCase
		"camelCase": func(val string) string {
			var (
				result = ""
				runes  = []rune(val)
				i      = 0
				upper  = false
			)

			for i < len(runes) {
				if runes[i] == '_' {
					i++
					upper = true
					continue
				}

				char := runes[i]
				if upper && runes[i-1] == '_' {
					char = unicode.ToUpper(char)
				}

				result += string(char)
				i++
				upper = false
			}

			return result
		},
		"getType": func(col *column) string {
			var dataType string
			switch col.Type {
			case "TEXT":
				dataType = "string"
			case "INTEGER":
				dataType = "int64"
			case "REAL":
				dataType = "float64"
			case "BLOB":
				return "[]bytes"
			case "NULL":
				return "{}interface"
			}

			if !col.NotNull {
				return fmt.Sprintf("*%s", dataType)
			}

			return dataType
		},
		"columnNames": func(cols []*column) []string {
			var items []string
			for _, col := range cols {
				items = append(items, col.Name)
			}

			return items
		},
		"join": func(items []string, sep string) string {
			var result string
			for i, item := range items {
				if i > 0 {
					result += sep
				}

				result += item
			}

			return result
		},
		"filter": func(items []string, excluded ...string) []string {
			var result []string
			for _, item := range items {
				match := false
				for _, ex := range excluded {
					if ex == item {
						match = true
					}
				}

				if !match {
					result = append(result, item)
				}
			}

			return result
		},
		"backtick": func() string {
			return "`"
		},
	}

	// Step 5) Execute template
	var buf bytes.Buffer
	tmpl := template.Must(template.New("").Funcs(funcs).Parse(genTmpl))
	err = tmpl.Execute(&buf, schemas)
	catch(err)

	// Step 6: Validate code

	// Step 7: Format code
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

// foo bar
var genTmpl = `
{{- print "package models" }}

// DO NOT EDIT! THIS IS GENERATED CODE!

{{ range . -}}
type {{ pascalCase .Name }} struct {
    {{ range .Columns }}
		{{ pascalCase .Name }} {{ getType . }} {{ jsonTag . -}}
    {{ end -}}
}

func Get{{ pascalCase .Name }}(db *sql.DB) ([]*{{ pascalCase .Name }}, error) {
	var items []*{{ pascalCase .Name }}
	query := {{ backtick }}
		select {{ join (columnNames .Columns) ", " }}
		from {{ .Name }};
	{{ backtick }}

	rows, err := db.Query(query)
	if err != nil {
		return items, err
	}

	defer func(rows *sql.Rows){
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	for rows.Next() {
		var item {{ pascalCase .Name }}		

		err = rows.Scan(
			&item.{{ range $index, $item := (columnNames .Columns) }}
			{{- if gt $index 0 }},
			&item.{{ end }}{{ pascalCase $item }}
			{{- end }},
		)

		if err != nil {
			return items, err
		}

		items = append(items, &item)
	}

	return items, nil
}
 
func Get{{ pascalCase .Name }}ById(db *sql.DB, id int64) (*{{ pascalCase .Name }}, error) {
	var item {{ pascalCase .Name }}
	query := {{ backtick }}
		select {{ join (columnNames .Columns) ", " }}
		from {{ .Name }}
		where id = ?;
	{{ backtick }}

	err := db.QueryRow(query, id).Scan(
		&item.{{ range $index, $col := .Columns }}
		{{- if gt $index 0 }},
		&item.{{ end }}{{ pascalCase $col.Name }}
		{{- end }},
	)

	if err != nil {
		return &item, err
	}

	return &item, nil
}

func Insert{{ pascalCase .Name }}(db *sql.DB, item *{{ pascalCase .Name }}) (*{{ pascalCase .Name }}, error) {
	query := {{ backtick }}
		insert into {{ .Name }}({{ join (filter (columnNames .Columns) "id") ", " }})
		values ({{ range $index, $col := (filter (columnNames .Columns) "id" "created_at" "updated_at") }}
		{{- if gt $index 0 }}, {{ end }}?
		{{- end }}, datetime(), datetime());
	{{ backtick }}

	result, err := db.Exec(
		query,
		&item.{{ range $index, $col := (filter (columnNames .Columns) "id" "created_at" "updated_at") }}
		{{- if gt $index 0 }},
		&item.{{ end }}{{ pascalCase $col }}
		{{- end }},
	)

	if err != nil {
		return item, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return item, err
	}

	return Get{{ pascalCase .Name }}ById(db, id)
}

func Update{{ pascalCase .Name }}(db *sql.DB, item *{{ pascalCase .Name }}) (*{{ pascalCase .Name }}, error) {
	query := {{ backtick }}
		update {{ .Name }}
		set {{ join (filter (columnNames .Columns ) "id" "created_at" "updated_at") " = ?, "}} = ?, updated_at = datetime()
	{{ backtick }}

	_, err := db.Exec(
		query,
		&item.{{ range $index, $col := (filter (columnNames .Columns) "id" "created_at" "updated_at") }}
		{{- if gt $index 0 }},
		&item.{{ end }}{{ pascalCase $col }}
		{{- end }},
	)

	if err != nil {
		return item, err
	}

	return Get{{ pascalCase .Name }}ById(db, item.Id)
}

func Delete{{ pascalCase .Name }}(db *sql.DB, id int) error {
	query := {{ backtick }}
		delete from {{ .Name }}
		where id = ?
	{{ backtick }}

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

{{ end -}}
	`
