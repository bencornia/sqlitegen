//go:generate go run github.com/bencornia/sqlitegen/cmd/sqlitegen -package-name basic -output-file basic.go basic.sqlite
package basic

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

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
}

func TestUpdateById(t *testing.T) {
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

			preUpdateItem, err := basicStore.GetById(t.Context(), itemId)
			if err != nil {
				t.Fatalf("failed to get item: %s", err.Error())
			}

			// Add a small delay so that we can see that the updated at field changed
			// updated_at will be the only field that should change since there are no extraneous columns
			time.Sleep(1 * time.Second)

			err = basicStore.UpdateById(t.Context(), preUpdateItem)
			if err != nil {
				t.Fatalf("failed to update item: %s", err.Error())
			}

			postUpdateItem, err := basicStore.GetById(t.Context(), itemId)
			if err != nil {
				t.Fatalf("failed to get item: %s", err.Error())
			}

			if postUpdateItem.Id != tc.expected.Id {
				t.Fatalf("expected id to equal %d but got %d", tc.expected.Id, postUpdateItem.Id)
			}

			if strings.Compare(preUpdateItem.UpdatedAt, postUpdateItem.UpdatedAt) == 0 {
				t.Errorf("expected updated_at to NOT be %s but got %s", preUpdateItem.UpdatedAt, postUpdateItem.UpdatedAt)
			}

			if strings.Compare(preUpdateItem.CreatedAt, postUpdateItem.CreatedAt) != 0 {
				t.Errorf("expected created_at to be %s but got %s", preUpdateItem.CreatedAt, postUpdateItem.CreatedAt)
			}
		})
	}
}

func TestDeleteById(t *testing.T) {
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

			err = basicStore.DeleteById(t.Context(), itemId)
			if err != nil {
				t.Fatalf("failed to delete item: %s", err.Error())
			}

			_, err = basicStore.GetById(t.Context(), itemId)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					t.Errorf("expected no rows to be returned but got: %s", err.Error())
				}
			}
		})
	}
}

func TestGetMany(t *testing.T) {
	// TODO:
}

func TestGetManyTx(t *testing.T) {
	// TODO:
}

func TestInsertManyTx(t *testing.T) {
	// TODO:
}

func TestUpdateManyTx(t *testing.T) {
	// TODO:
}

func TestDeleteManyTx(t *testing.T) {
	testCases := []struct {
		name          string
		items         []*Basic
		ids           []int64 // Id's of items we want to delete
		expectedCount int64
	}{
		{
			name:          "OneItemDeleteAll",
			items:         []*Basic{{}},
			ids:           []int64{1},
			expectedCount: 0,
		},
		{
			name:          "TwoItemsDeleteAll",
			items:         []*Basic{{}, {}},
			ids:           []int64{1, 2},
			expectedCount: 0,
		},
		{
			name:          "TwoItemsDeleteOne",
			items:         []*Basic{{}, {}},
			ids:           []int64{1},
			expectedCount: 1,
		},
		{
			name:          "MultipleItemsDeleteAll",
			items:         []*Basic{{}, {}, {}},
			ids:           []int64{1, 2, 3},
			expectedCount: 0,
		},
		{
			name:          "MultipleItemsDeleteOne",
			items:         []*Basic{{}, {}, {}},
			ids:           []int64{2},
			expectedCount: 2,
		},
	}

	sqlStmt := getSQLString(t, "basic.sql")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := setupTmpDB(t, "basic.sqlite", sqlStmt)
			defer db.Close()

			basicStore := NewBasicStore(db)
			ctx := t.Context()
			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				t.Fatalf("failed to create transaction: %s", err.Error())
			}

			_, err = basicStore.InsertManyTx(ctx, tx, tc.items)
			if err != nil {
				t.Fatalf("failed to insert items: %s", err.Error())
			}

			err = basicStore.DeleteManyTx(ctx, tx, tc.ids)
			if err != nil {
				t.Fatalf("failed to delete items: %s", err.Error())
			}

			for _, id := range tc.ids {
				_, err = basicStore.GetByIdTx(ctx, tx, id)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						t.Errorf("expected no rows to be returned but got: %s", err.Error())
					}
				} else {
					t.Errorf("expected an error")
				}
			}
		})
	}
}
