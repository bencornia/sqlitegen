package codegen

import (
	"database/sql"
	"strings"
	"testing"
)

func TestPascalCase(t *testing.T) {
	testCases := []struct {
		name   string
		given  string
		expect string
	}{
		{name: "EmptyString", given: "", expect: ""},
		{name: "Basic", given: "foo", expect: "Foo"},
		{name: "AllCaps", given: "FOO", expect: "Foo"},
		{name: "PascalCase", given: "FooBar", expect: "Foobar"},
		{name: "SnakeCase", given: "foo_bar", expect: "FooBar"},
		{name: "SnakeCaseThreeWords", given: "foo_bar_baz", expect: "FooBarBaz"},
		{name: "LeadingSnakeCase", given: "_foo_bar", expect: "FooBar"},
		{name: "TrailingSnakeCase", given: "foo_bar_", expect: "FooBar"},
		{name: "WrappedSnakeCase", given: "_foo_bar_", expect: "FooBar"},
		{name: "MultipleUnderscores", given: "foo__bar", expect: "FooBar"},
		{name: "MultipleTrailingUnderscores", given: "foo__bar__", expect: "FooBar"},
		{name: "SingleLetter", given: "f", expect: "F"},
		{name: "LeadingUnderscoreSingleLetter", given: "_f", expect: "F"},
		{name: "SingleLetterSnakeCase", given: "f_b", expect: "FB"},
		{name: "SingleUnderscore", given: "_", expect: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := pascalCase(tc.given)
			if strings.Compare(actual, tc.expect) != 0 {
				t.Errorf("given [%s] expect [%s] but got [%s]", tc.given, tc.expect, actual)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	testCases := []struct {
		name   string
		given  string
		expect string
	}{
		{name: "EmptyString", given: "", expect: ""},
		{name: "SingleWord", given: "foo", expect: "foo"},
		{name: "MixedCaps", given: "fOo", expect: "foo"},
		{name: "SnakeCase", given: "foo_bar", expect: "fooBar"},
		{name: "SnakeCaseThreeWords", given: "foo_bar_baz", expect: "fooBarBaz"},
		{name: "MultipleUnderscores", given: "foo__bar", expect: "fooBar"},
		{name: "TrailingSnakeCase", given: "foo_bar_", expect: "fooBar"},
		{name: "WrappedSnakeCase", given: "_foo_bar_", expect: "fooBar"},
		{name: "LeadingSnakeCase", given: "_foo_bar", expect: "fooBar"},
		{name: "SingleUnderScore", given: "_", expect: ""},
		{name: "LeadingUnderscoreSingleLetter", given: "_f", expect: "f"},
		{name: "SingleLetterSnakeCase", given: "f_b", expect: "fB"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := camelCase(tc.given)
			if strings.Compare(actual, tc.expect) != 0 {
				t.Errorf("given [%s] expect [%s] but got [%s]", tc.given, tc.expect, actual)
			}
		})
	}
}

func TestJoinItems(t *testing.T) {
	testCases := []struct {
		name   string
		given  []string
		sep    string
		expect string
	}{
		{
			name:   "EmptyStringJoinedByComma",
			given:  []string{},
			sep:    ", ",
			expect: "",
		},
		{
			name:   "SingleWordJoinedByComma",
			given:  []string{"foo"},
			sep:    ", ",
			expect: "foo",
		},
		{
			name:   "TwoWordsJoinedByComma",
			given:  []string{"foo", "bar"},
			sep:    ", ",
			expect: "foo, bar",
		},
		{
			name:   "ThreeWordsJoinedByComma",
			given:  []string{"foo", "bar", "baz"},
			sep:    ", ",
			expect: "foo, bar, baz",
		},
		{
			name:   "ThreeWordsJoinedByOr",
			given:  []string{"foo", "bar", "baz"},
			sep:    " || ",
			expect: "foo || bar || baz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := joinItems(tc.given, tc.sep)
			if strings.Compare(actual, tc.expect) != 0 {
				t.Errorf("given [%s] expect [%s] but got [%s]", tc.given, tc.expect, actual)
			}
		})
	}
}

func TestColumnNames(t *testing.T) {
	testCases := []struct {
		name   string
		given  []*column
		expect []string
	}{
		{
			name:   "NoColumns",
			given:  []*column{},
			expect: []string{},
		},
		{
			name:   "OneColumn",
			given:  []*column{{Name: "foo"}},
			expect: []string{"foo"},
		},
		{
			name:   "TwoColumns",
			given:  []*column{{Name: "foo"}, {Name: "bar"}},
			expect: []string{"foo", "bar"},
		},
		{
			name:   "ThreeColumms",
			given:  []*column{{Name: "foo"}, {Name: "bar"}, {Name: "baz"}},
			expect: []string{"foo", "bar", "baz"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := columnNames(tc.given)
			if len(actual) != len(tc.expect) {
				t.Fatalf("length of actual [%d] does not match length of given [%d]", len(actual), len(tc.expect))
			}

			for i := range actual {
				if strings.Compare(actual[i], tc.expect[i]) != 0 {
					t.Errorf("expect [%s] but got [%s]", tc.expect[i], actual[i])
				}
			}
		})
	}
}

func TestMapItems(t *testing.T) {
	testCases := []struct {
		name   string
		given  []string
		val    string
		expect []string
	}{
		{
			name:   "EmptyItems",
			given:  []string{},
			val:    "",
			expect: []string{},
		},
		{
			name:   "SingleWordReplacedByQuestionMark",
			given:  []string{"foo"},
			val:    "?",
			expect: []string{"?"},
		},
		{
			name:   "TwoWordsReplacedByQuestionMark",
			given:  []string{"foo", "bar"},
			val:    "?",
			expect: []string{"?", "?"},
		},
		{
			name:   "ThreeWordsReplacedByQuestionMark",
			given:  []string{"foo", "bar", "baz"},
			val:    "?",
			expect: []string{"?", "?", "?"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := mapItems(tc.given, tc.val)
			if len(actual) != len(tc.expect) {
				t.Fatalf("length of actual [%d] does not match length of expected [%d]", len(actual), len(tc.expect))
			}
			for i := range actual {
				if strings.Compare(actual[i], tc.expect[i]) != 0 {
					t.Errorf("expect [%s] but got [%s]", tc.expect[i], actual[i])
				}
			}
		})
	}
}

func TestFilterItems(t *testing.T) {
	testCases := []struct {
		name     string
		given    []string
		excluded []string
		expect   []string
	}{
		{
			name:     "EmptyItems",
			given:    []string{},
			excluded: []string{},
			expect:   []string{},
		},
		{
			name:     "SingleItemNoFilter",
			given:    []string{"foo"},
			excluded: []string{},
			expect:   []string{"foo"},
		},
		{
			name:     "SingleItemSingleFilter",
			given:    []string{"foo"},
			excluded: []string{"foo"},
			expect:   []string{},
		},
		{
			name:     "TwoItemsOneFilter",
			given:    []string{"foo", "bar"},
			excluded: []string{"foo"},
			expect:   []string{"bar"},
		},
		{
			name:     "ThreeItemsOneFilter",
			given:    []string{"foo", "bar", "baz"},
			excluded: []string{"bar"},
			expect:   []string{"foo", "baz"},
		},
		{
			name:     "ThreeItemsTwoFilters",
			given:    []string{"foo", "bar", "baz"},
			excluded: []string{"bar", "baz"},
			expect:   []string{"foo"},
		},
		{
			name:     "ThreeItemsThreeFilters",
			given:    []string{"foo", "bar", "baz"},
			excluded: []string{"bar", "baz", "foo"},
			expect:   []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := filterItems(tc.given, tc.excluded...)
			if len(actual) != len(tc.expect) {
				t.Fatalf("length of actual [%d] does not match length of expected [%d]", len(actual), len(tc.expect))
			}
			for i := range actual {
				if strings.Compare(actual[i], tc.expect[i]) != 0 {
					t.Errorf("expect [%s] but got [%s]", tc.expect[i], actual[i])
				}
			}
		})
	}
}

func TestBacktick(t *testing.T) {
	if strings.Compare("`", backtick()) != 0 {
		t.Errorf("backtick() does not return `")
	}
}

func TestGetType(t *testing.T) {
	testCases := []struct {
		name     string
		sql      string
		expected string
	}{
		{
			name:     "IntType",
			sql:      "create table foo(a int not null)",
			expected: "int64",
		},
		{
			name:     "IntegerType",
			sql:      "create table foo(a integer not null)",
			expected: "int64",
		},
		{
			name:     "RealType",
			sql:      "create table foo(a real not null)",
			expected: "float64",
		},
		{
			name:     "TextType",
			sql:      "create table foo(a text not null)",
			expected: "string",
		},
		{
			name:     "BlobType",
			sql:      "create table foo(a blob not null)",
			expected: "[]bytes",
		},
		{
			name:     "AnyType",
			sql:      "create table foo(a any not null)",
			expected: "any",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, err := sql.Open("sqlite3", ":memory:")
			if err != nil {
				t.Fatalf("Failed to open database: %s", err.Error())
			}

			defer catchClosable(db)

			_, err = db.Exec(tc.sql)
			if err != nil {
				t.Fatalf("Failed to create table: %s", err.Error())
			}

			columns, err := getColumns(db, "foo")
			if err != nil {
				t.Fatalf("Failed to get columns: %s", err.Error())
			}

			if len(columns) != 1 {
				t.Fatalf("testcase should have exactly one column")
			}

			col := columns[0]

			// TODO: I think we should assert that a panic happens
			if !col.NotNull {
				t.Fatalf("column must be not null")
			}

			colType := getType(col)
			if strings.Compare(colType, tc.expected) != 0 {
				t.Errorf("expected [%s] but got [%s]", tc.expected, colType)
			}
		})
	}
}
