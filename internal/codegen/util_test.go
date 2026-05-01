package codegen

import (
	"database/sql"
	"strings"
	"testing"
)

func TestPascalCase(t *testing.T) {
	testCases := []struct {
		given  string
		expect string
	}{
		{given: "", expect: ""},
		{given: "foo", expect: "Foo"},
		{given: "FOO", expect: "Foo"},
		{given: "FooBar", expect: "Foobar"},
		{given: "foo_bar", expect: "FooBar"},
		{given: "foo_bar_baz", expect: "FooBarBaz"},
		{given: "_foo_bar", expect: "FooBar"},
		{given: "foo_bar_", expect: "FooBar"},
		{given: "_foo_bar_", expect: "FooBar"},
		{given: "foo__bar", expect: "FooBar"},
		{given: "foo__bar__", expect: "FooBar"},
		{given: "f", expect: "F"},
		{given: "_f", expect: "F"},
		{given: "f_b", expect: "FB"},
		{given: "_", expect: ""},
	}

	for _, tc := range testCases {
		actual := pascalCase(tc.given)
		if strings.Compare(actual, tc.expect) != 0 {
			t.Errorf("given [%s] expect [%s] but got [%s]", tc.given, tc.expect, actual)
		}
	}
}

func TestCamelCase(t *testing.T) {
	testCases := []struct {
		given  string
		expect string
	}{
		{given: "", expect: ""},
		{given: "foo", expect: "foo"},
		{given: "fOo", expect: "foo"},
		{given: "foo_bar", expect: "fooBar"},
		{given: "foo_bar_baz", expect: "fooBarBaz"},
		{given: "foo__bar", expect: "fooBar"},
		{given: "foo_bar_", expect: "fooBar"},
		{given: "_foo_bar_", expect: "fooBar"},
		{given: "_foo_bar", expect: "fooBar"},
		{given: "_", expect: ""},
		{given: "_f", expect: "f"},
		{given: "f_b", expect: "fB"},
	}

	for _, tc := range testCases {
		actual := camelCase(tc.given)
		if strings.Compare(actual, tc.expect) != 0 {
			t.Errorf("given [%s] expect [%s] but got [%s]", tc.given, tc.expect, actual)
		}
	}
}

func TestJoinItems(t *testing.T) {
	testCases := []struct {
		given  []string
		sep    string
		expect string
	}{
		{
			given:  []string{},
			sep:    ", ",
			expect: "",
		},
		{
			given:  []string{"foo"},
			sep:    ", ",
			expect: "foo",
		},
		{
			given:  []string{"foo", "bar"},
			sep:    ", ",
			expect: "foo, bar",
		},
		{
			given:  []string{"foo", "bar", "baz"},
			sep:    ", ",
			expect: "foo, bar, baz",
		},
		{
			given:  []string{"foo", "bar", "baz"},
			sep:    " || ",
			expect: "foo || bar || baz",
		},
	}

	for _, tc := range testCases {
		actual := joinItems(tc.given, tc.sep)
		if strings.Compare(actual, tc.expect) != 0 {
			t.Errorf("given [%s] expect [%s] but got [%s]", tc.given, tc.expect, actual)
		}
	}
}

func TestColumnNames(t *testing.T) {
	testCases := []struct {
		given  []*column
		expect []string
	}{
		{
			given:  []*column{},
			expect: []string{},
		},
		{
			given:  []*column{{Name: "foo"}},
			expect: []string{"foo"},
		},
		{
			given:  []*column{{Name: "foo"}, {Name: "bar"}},
			expect: []string{"foo", "bar"},
		},
		{
			given:  []*column{{Name: "foo"}, {Name: "bar"}, {Name: "baz"}},
			expect: []string{"foo", "bar", "baz"},
		},
	}

	for _, tc := range testCases {
		actual := columnNames(tc.given)
		if len(actual) != len(tc.expect) {
			t.Fatalf("length of actual [%d] does not match length of given [%d]", len(actual), len(tc.expect))
		}

		for i := range actual {
			if strings.Compare(actual[i], tc.expect[i]) != 0 {
				t.Errorf("expect [%s] but got [%s]", tc.expect[i], actual[i])
			}
		}
	}
}

func TestMapItems(t *testing.T) {
	testCases := []struct {
		given  []string
		val    string
		expect []string
	}{
		{
			given:  []string{},
			val:    "",
			expect: []string{},
		},
		{
			given:  []string{"foo"},
			val:    "?",
			expect: []string{"?"},
		},
		{
			given:  []string{"foo", "bar"},
			val:    "?",
			expect: []string{"?", "?"},
		},
		{
			given:  []string{"foo", "bar", "baz"},
			val:    "?",
			expect: []string{"?", "?", "?"},
		},
	}

	for _, tc := range testCases {
		actual := mapItems(tc.given, tc.val)
		if len(actual) != len(tc.expect) {
			t.Fatalf("length of actual [%d] does not match length of expected [%d]", len(actual), len(tc.expect))
		}
		for i := range actual {
			if strings.Compare(actual[i], tc.expect[i]) != 0 {
				t.Errorf("expect [%s] but got [%s]", tc.expect[i], actual[i])
			}
		}
	}
}

func TestFilterItems(t *testing.T) {
	testCases := []struct {
		given    []string
		excluded []string
		expect   []string
	}{
		{
			given:    []string{},
			excluded: []string{},
			expect:   []string{},
		},
		{
			given:    []string{"foo"},
			excluded: []string{},
			expect:   []string{"foo"},
		},
		{
			given:    []string{"foo"},
			excluded: []string{"foo"},
			expect:   []string{},
		},
		{
			given:    []string{"foo", "bar"},
			excluded: []string{"foo"},
			expect:   []string{"bar"},
		},
		{
			given:    []string{"foo", "bar", "baz"},
			excluded: []string{"bar"},
			expect:   []string{"foo", "baz"},
		},
		{
			given:    []string{"foo", "bar", "baz"},
			excluded: []string{"bar", "baz"},
			expect:   []string{"foo"},
		},
		{
			given:    []string{"foo", "bar", "baz"},
			excluded: []string{"bar", "baz", "foo"},
			expect:   []string{},
		},
	}

	for _, tc := range testCases {
		actual := filterItems(tc.given, tc.excluded...)
		if len(actual) != len(tc.expect) {
			t.Fatalf("length of actual [%d] does not match length of expected [%d]", len(actual), len(tc.expect))
		}
		for i := range actual {
			if strings.Compare(actual[i], tc.expect[i]) != 0 {
				t.Errorf("expect [%s] but got [%s]", tc.expect[i], actual[i])
			}
		}
	}
}

func TestBacktick(t *testing.T) {
	if strings.Compare("`", backtick()) != 0 {
		t.Errorf("backtick() does not return `")
	}
}

func TestGetType(t *testing.T) {
	// TODO: Create in memory sqlite connection
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open sqlite in memory connection")
	}

	_, err = db.Exec("create table employee(id integer primary key, name text, salary real, profile_photo blob, )")
	if err != nil {
		t.Fatalf("Failed to create table")
	}

	// TODO: a little tricky here with sqlite's loose typing
}
