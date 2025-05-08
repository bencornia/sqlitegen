package codegen

import (
	"database/sql"
	"fmt"
	"log"

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

func Generate(dbPath string, outFile string) {
	db := newDB(dbPath)

	var exists bool
	err := db.QueryRow("select count(*) > 0 from sqlite_master").Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.Fatal("No tables!")
	}

	// type|name|tbl_name|rootpage|sql
	tableNames, err := db.Query("select name from sqlite_master where type = 'table'")
	if err != nil {
		log.Fatal(err)
	}

	var schemas []*schema
	defer closeRows(tableNames)
	for tableNames.Next() {
		var tableName string
		err = tableNames.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}

		c, err := db.Query("select name, type, `notnull` from pragma_table_info('?')", tableName)
		if err != nil {
			log.Fatal(err)
		}

		defer closeRows(c)
		s := &schema{Name: tableName, Columns: []*column{}}
		for c.Next() {
			var col column
			err = c.Scan(&col.Name, &col.Type, &col.NotNull)
			if err != nil {
				log.Fatal(err)
			}

			s.Columns = append(s.Columns, &col)
		}

		schemas = append(schemas, s)
	}

	fmt.Println(schemas)
}

func newDB(dsn string) *sql.DB {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func closeRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}
