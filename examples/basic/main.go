//go:generate go run github.com/bencornia/sqlitegen/cmd/sqlitegen -output internal/model/models.go db.sqlite
package main

import (
	"database/sql"
	"fmt"

	"github.com/bencornia/sqlitegen/examples/basic/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")

	store := model.NewStore(db)
	employee := &model.Employee{
		FirstName: "Scottie",
		LastName:  "Andrus",
		Age:       64,
	}

	employee, _ = store.InsertEmployee(employee)
	fmt.Println(employee)

	employee.Age = 62
	employee, _ = store.UpdateEmployee(employee)
	fmt.Println(employee)

	_ = store.DeleteEmployee(int(employee.Id))
}
