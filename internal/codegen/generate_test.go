package codegen

import (
	"database/sql"
	"strings"
	"testing"
)

func TestIsValidSchema(t *testing.T) {
	testCases := []struct {
		name       string
		expected   bool
		tableName  string
		createStmt string
	}{
		{
			name:      "IsNotStrict",
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
			name:      "IsValid",
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
			name:      "IsMissingPrimaryKey",
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
			name:      "IsMissingNotNullConstraint",
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
			name:      "IsMissingUpdatedAt",
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
			name:      "IsMissingCreatedAt",
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
		t.Run(tc.name, func(t *testing.T) {
			db, err := sql.Open("sqlite3", ":memory:")
			if err != nil {
				t.Fatal("Failed to open database")
			}

			defer catchClosable(db)

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
		})
	}
}

func TestGetSchemas(t *testing.T) {
	testCases := []struct {
		name    string
		sql     string
		schemas []*schema
	}{
		{
			name:    "NoSchemas",
			sql:     ``,
			schemas: []*schema{},
		},
		{
			name: "IsValid",
			sql: `
				create table foo(
					id integer primary key not null,
					updated_at text not null,
					created_at text not null
				) strict;
			`,
			schemas: []*schema{
				{
					Name: "foo",
					Columns: []*column{
						{
							Name:         "id",
							NotNull:      true,
							IsPrimaryKey: true,
							Type:         "INTEGER",
						},
						{
							Name:         "updated_at",
							NotNull:      true,
							IsPrimaryKey: false,
							Type:         "TEXT",
						},
						{
							Name:         "created_at",
							NotNull:      true,
							IsPrimaryKey: false,
							Type:         "TEXT",
						},
					},
				},
			},
		},
		{
			name: "OneInvalidSchema",
			sql: `
				create table foo(
					id integer primary key not null,
					updated_at text not null,
					created_at text not null
				);
			`,
			schemas: []*schema{},
		},
		{
			name: "MixedValidity",
			sql: `
				create table foo(
					id integer primary key not null,
					updated_at text not null,
					created_at text not null
				) strict;

				create table bar(
					id integer primary key not null,
					updated_at text not null,
					created_at text not null
				);
			`,
			schemas: []*schema{
				{
					Name: "foo",
					Columns: []*column{
						{
							Name:         "id",
							NotNull:      true,
							IsPrimaryKey: true,
							Type:         "INTEGER",
						},
						{
							Name:         "updated_at",
							NotNull:      true,
							IsPrimaryKey: false,
							Type:         "TEXT",
						},
						{
							Name:         "created_at",
							NotNull:      true,
							IsPrimaryKey: false,
							Type:         "TEXT",
						},
					},
				},
			},
		},
	}

	// TODO: Add tests for these helpers???
	columnDeepEqual := func(a *column, b *column) bool {
		return a.IsPrimaryKey == b.IsPrimaryKey &&
			strings.Compare(a.Name, b.Name) == 0 &&
			a.NotNull == b.NotNull &&
			strings.Compare(a.Type, b.Type) == 0
	}

	schemaDeepEqual := func(a *schema, b *schema) bool {
		if len(a.Columns) != len(b.Columns) {
			return false
		}

		columnsEqual := true
		for i := range len(a.Columns) {
			columnsEqual = columnsEqual && columnDeepEqual(a.Columns[i], b.Columns[i])
		}

		return columnsEqual
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, err := sql.Open("sqlite3", ":memory:")
			if err != nil {
				t.Fatalf("failed to open in memory database: %s", err.Error())
			}

			defer catchClosable(db)

			_, err = db.Exec(tc.sql)
			if err != nil {
				t.Fatalf("failed to execute query: %s", err.Error())
			}

			schemas, err := getSchemas(db)
			if err != nil {
				t.Fatalf("failed to getSchemas: %s", err.Error())
			}

			// Compare schemas
			if len(schemas) != len(tc.schemas) {
				t.Fatalf("expected len(schemas) to equal %d but got %d", len(tc.schemas), len(schemas))
			}

			for i := range len(schemas) {
				if !schemaDeepEqual(schemas[i], tc.schemas[i]) {
					t.Errorf("expected schema for [%s] does not match actual", tc.schemas[i].Name)
				}
			}
		})
	}
}

func TestGetColumns(t *testing.T) {
	// TODO:
}

func TestGetTableNames(t *testing.T) {
	// TODO:
}
