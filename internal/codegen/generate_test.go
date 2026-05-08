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

func columnDeepEqual(t *testing.T, a *column, b *column) bool {
	t.Helper()
	return a.IsPrimaryKey == b.IsPrimaryKey &&
		strings.Compare(a.Name, b.Name) == 0 &&
		a.NotNull == b.NotNull &&
		strings.Compare(a.Type, b.Type) == 0
}

func schemaDeepEqual(t *testing.T, a *schema, b *schema) bool {
	t.Helper()
	if len(a.Columns) != len(b.Columns) {
		return false
	}

	columnsEqual := true
	for i := range len(a.Columns) {
		columnsEqual = columnsEqual && columnDeepEqual(t, a.Columns[i], b.Columns[i])
	}

	return columnsEqual
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
				if !schemaDeepEqual(t, schemas[i], tc.schemas[i]) {
					t.Errorf("expected schema for [%s] does not match actual", tc.schemas[i].Name)
				}
			}
		})
	}
}

func TestGetType(t *testing.T) {
	testCases := []struct {
		name     string
		col      column
		expected string
	}{
		{
			name:     "IntType",
			col:      column{Name: "foo", Type: "INT", NotNull: true, IsPrimaryKey: false},
			expected: "int64",
		},
		{
			name:     "IntegerType",
			col:      column{Name: "foo", Type: "INTEGER", NotNull: true, IsPrimaryKey: false},
			expected: "int64",
		},
		{
			name:     "RealType",
			col:      column{Name: "foo", Type: "REAL", NotNull: true, IsPrimaryKey: false},
			expected: "float64",
		},
		{
			name:     "TextType",
			col:      column{Name: "foo", Type: "TEXT", NotNull: true, IsPrimaryKey: false},
			expected: "string",
		},
		{
			name:     "BlobType",
			col:      column{Name: "foo", Type: "BLOB", NotNull: true, IsPrimaryKey: false},
			expected: "[]bytes",
		},
		{
			name:     "AnyType",
			col:      column{Name: "foo", Type: "ANY", NotNull: true, IsPrimaryKey: false},
			expected: "any",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			colType := getType(&tc.col)
			if strings.Compare(colType, tc.expected) != 0 {
				t.Errorf("expected [%s] but got [%s]", tc.expected, colType)
			}
		})
	}
}

func TestGetTypePanics(t *testing.T) {
	testCases := []struct {
		name     string
		col      column
		expected string
	}{
		{
			name:     "NullColumnPanics",
			col:      column{Name: "foo", Type: "TEXT", NotNull: false, IsPrimaryKey: false},
			expected: "columns must be not null",
		},
		{
			name:     "UnknownTypePanics",
			col:      column{Name: "foo", Type: "SOME_UNKNOWN_TYPE", NotNull: true, IsPrimaryKey: false},
			expected: "unknown datatype",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				tc := tc
				r := recover()
				if r == nil {
					t.Fatalf("expected panic but none occurred")
				}

				msg, ok := r.(string)
				if !ok {
					t.Fatalf("an unkown panic occurred: %v", r)
				}

				if strings.Compare(msg, tc.expected) != 0 {
					t.Errorf("expected [%s] but got [%s]", tc.expected, msg)
				}
			}()
			getType(&tc.col)
		})
	}
}
