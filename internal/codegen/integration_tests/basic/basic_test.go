//go:generate go run github.com/bencornia/sqlitegen/cmd/sqlitegen -package-name basic -output-file basic.go basic.sqlite
package basic

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func getSQLString(t *testing.T, filename string) string {
	t.Helper()

	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open basic.sql: %s", err.Error())
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("failed to read basic.sql: %s", err.Error())
	}

	return string(buf)
}

func setupTmpDB(t *testing.T, dbFilename string, sqlStmt string) *sql.DB {
	t.Helper()
	tmpDB := fmt.Sprintf("%s/%s", t.TempDir(), dbFilename)
	db, err := sql.Open("sqlite3", tmpDB)
	if err != nil {
		t.Fatalf("failed to open database: %s", err.Error())
	}

	_, err = db.Exec(sqlStmt)
	if err != nil {
		t.Fatalf("failed to create table: %s", err.Error())
	}

	return db
}

func TestInsertAndGetById(t *testing.T) {
	testCases := []struct {
		name     string
		item     Basic
		expected Basic
	}{
		{
			name:     "SimpleItem",
			item:     Basic{},
			expected: Basic{Id: 1},
		},
	}

	sqlStmt := getSQLString(t, "basic.sql")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := setupTmpDB(t, "basic.sqlite", sqlStmt)
			defer db.Close()

			basicStore := NewBasicStore(db)
			itemId, err := basicStore.Insert(t.Context(), &tc.item)
			if err != nil {
				t.Fatalf("failed to insert item: %s", err.Error())
			}

			item, err := basicStore.GetById(t.Context(), itemId)
			if err != nil {
				t.Fatalf("failed to get item: %s", err.Error())
			}

			if item.Id != tc.expected.Id {
				t.Errorf("expected id to equal %d but got %d", tc.expected.Id, item.Id)
			}
		})
	}

	func TestInsertAndUpdateAndGet(t *testing.T) {
		
	}
}
