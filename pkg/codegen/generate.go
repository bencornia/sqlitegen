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
			return "any"
		},
	}

	// Step 5) Execute template
	var buf bytes.Buffer
	tmpl := template.Must(template.New("").Funcs(funcs).Parse(genTmpl))
	err = tmpl.Execute(&buf, schemas)
	catch(err)

	// Step 6: Validate code

	// Step 7: Format code
	formatted, err := imports.Process("foo.go", buf.Bytes(), nil)
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
type {{ .Name | pascalCase }} struct {
    {{ range .Columns }}
        {{ .Name | pascalCase }} {{ . | getType }} {{ . | jsonTag }}
    {{ end -}}
}

func GetAll{{ .Name | pascalCase }}() ([]*{{ .Name | pascalCase }}, error) {
    
}

func Get{{ .Name | pascalCase }}(id int) (*{{ .Name | pascalCase }}, error) {
    
}

func Update{{ .Name | pascalCase }}({{ .Name }} *{{ .Name | pascalCase }}) (*{{ .Name | pascalCase }}, error) {
    
}

func Delete{{ .Name | pascalCase }}(id int) (error) {
    
}
{{ end -}}
	`
