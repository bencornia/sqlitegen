package codegen

import (
	"database/sql"
	"testing"
)

func TestIsValidSchema(t *testing.T) {
	testCases := []struct {
		expected   bool
		tableName  string
		createStmt string
	}{
		{
			expected:  false,
			tableName: "not_strict",
			createStmt: `
				create table not_strict (
					id integer primary key not null,	
					created_at text not null,
					updated_at text not null
				);
			`,
		},
		{
			expected:  true,
			tableName: "is_valid",
			createStmt: `
				create table is_valid(
					id integer primary key not null,	
					created_at text not null,
					updated_at text not null
				) strict;
			`,
		},
		{
			expected:  false,
			tableName: "missing_pk",
			createStmt: `
				create table missing_pk (
					id integer not null,	
					created_at text not null,
					updated_at text not null
				) strict;
			`,
		},
		{
			expected:  false,
			tableName: "missing_not_null",
			createStmt: `
				create table missing_not_null (
					id integer primary key not null,	
					name text,
					created_at text not null,
					updated_at text not null
				) strict;
			`,
		},
		{
			expected:  false,
			tableName: "missing_updated_at",
			createStmt: `
				create table missing_updated_at (
					id integer primary key not null,	
					created_at text not null
				) strict;
			`,
		},
		{
			expected:  false,
			tableName: "missing_created_at",
			createStmt: `
				create table missing_created_at (
					id integer primary key not null,	
					updated_at text not null
				) strict;
			`,
		},
	}

	for _, tc := range testCases {
		func() {
			db, err := sql.Open("sqlite3", ":memory:")
			if err != nil {
				t.Fatal("Failed to open database")
			}

			defer db.Close()

			_, err = db.Exec(tc.createStmt)
			if err != nil {
				t.Fatalf("the create statement for [%s] failed with %s", tc.tableName, err.Error())
			}

			ok, err := isValidSchema(db, tc.tableName)
			if err != nil {
				t.Fatalf("validating %s failed with %s", tc.tableName, err.Error())
			}

			if ok != tc.expected {
				t.Errorf("expected %v got %v for %s", tc.expected, ok, tc.tableName)
			}
		}()
	}
}

func TestGetSchemas(t *testing.T) {
	testCases := []struct {
		sql      string
		expected []*column
	}{
		{
			sql:      ``,
			expected: []*column{},
		},
	}

	for _, tc := range testCases {
		func() {
			db, err := sql.Open("sqlite3", ":memory:")
			if err != nil {
				t.Fatalf("failed to open in memory database: %s", err.Error())
			}

			defer db.Close()

			_, err = db.Exec(tc.sql)
			if err != nil {
				t.Fatalf("failed to execute query: %s", err.Error())
			}

			schemas, err := getSchemas(db)
			if err != nil {
				t.Fatalf("failed to getSchemas: %s", err.Error())
			}

			if len(schemas) != len(tc.expected) {
				t.Fatalf("got len(schemas) %d but got %d", len(schemas), len(tc.expected))
			}

			// TODO: Do deep comparison of each column
		}()
	}
}
