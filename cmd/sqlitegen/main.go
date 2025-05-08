package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Schema struct {
	Name    string
	Columns []*Column
}

type Column struct {
	Name    string
	Type    string
	NotNull bool
}

func main() {
	db := newDB()

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

	var schemas []*Schema
	defer closeRows(tableNames)
	for tableNames.Next() {
		var tableName string
		err = tableNames.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}

		columns, err := db.Query("select name, type, `notnull` from pragma_table_info('?')", tableName)
		if err != nil {
			log.Fatal(err)
		}

		defer closeRows(columns)
		schema := &Schema{Name: tableName, Columns: []*Column{}}
		for columns.Next() {
			var col Column
			err = columns.Scan(&col.Name, &col.Type, &col.NotNull)
			if err != nil {
				log.Fatal(err)
			}

			schema.Columns = append(schema.Columns, &col)
		}

		schemas = append(schemas, schema)
	}

	fmt.Println(schemas)
}

func newDB() *sql.DB {
	var dsn string
	if dsn = os.Getenv("DB_PATH"); dsn == "" {
		log.Fatal("Missing DB_PATH")
	}

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
