package codegen

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"text/template"
	"unicode"

	_ "github.com/mattn/go-sqlite3"
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
		"title": func(val string) string {
			runes := []rune(val)
			runes[0] = unicode.ToUpper(runes[0])
			return string(runes)
		},
		"jsonTag": func(col *column) string {
			return fmt.Sprintf("`json:\"%s\"`", col.Name)
		},
	}

	// Step 5) Execute template
	tmpl := template.Must(template.New("").Funcs(funcs).Parse(genTmpl))
	err = tmpl.Execute(writer, schemas)
	catch(err)
}

var genTmpl = `
package models

{{ range . }}
	type {{ .Name | title }} struct {
		{{ range .Columns }}
			{{ .Name | title }} {{ . | jsonTag }}
		{{ end }}
	}

	func GetAll{{ .Name | title }}() ([]*{{ .Name | title }}, error) {
		
	}

	func Get{{ .Name | title }}(id int) (*{{ .Name | title }}, error) {
		
	}

	func Update{{ .Name | title }}({{ .Name }} *{{ .Name | title }}) (*{{ .Name | title }}, error) {
		
	}

	func Delete{{ .Name | title }}(id int) (error) {
		
	}
{{ end -}}
	`
