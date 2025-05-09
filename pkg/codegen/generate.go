package codegen

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"text/template"

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
		catch(fmt.Errorf("No tables in %s", dsn))
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

		c, err := db.Query("select name, type, `notnull` from pragma_table_info('?')", tableName)
		catch(err)

		defer func(rows *sql.Rows) {
			catch(rows.Close())
		}(c)

		s := &schema{Name: tableName, Columns: []*column{}}
		for c.Next() {
			var col column
			err = c.Scan(&col.Name, &col.Type, &col.NotNull)
			catch(err)

			s.Columns = append(s.Columns, &col)
		}

		schemas = append(schemas, s)
	}

	// Step 4) Execute template
	tmpl := template.Must(template.New("").Parse(genTmpl))
	err = tmpl.Execute(writer, schemas)
	catch(err)
}

var genTmpl = `
package models

{{ range . }}

{{ end }}
	`
