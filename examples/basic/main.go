//go:generate go run github.com/bencornia/sqlitegen/cmd/sqlitegen -output-file internal/model/models.go db.sqlite
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/bencornia/sqlitegen/examples/basic/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	employeeStore := model.NewEmployeeStore(db)
	ids, err := employeeStore.InsertMany(
		ctx,
		tx,
		[]*model.Employee{
			{Name: "George Burdell"},
			{Name: "Abraham Lincoln"},
			{Name: "Alexander Hamilton"},
		},
	)

	if err != nil {
		_ = tx.Rollback()
	}

	employees, err := employeeStore.GetManyTx(ctx, tx, ids)
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(employees)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
