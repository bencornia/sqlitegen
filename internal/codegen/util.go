package codegen

import (
	"fmt"
	"unicode"
)

func getTag(col *column) string {
	return fmt.Sprintf("`json:\"%s\"`", col.Name)
}

func pascalCase(val string) string {
	var (
		result = ""
		runes  = []rune(val)
		i      = 0
		upper  = false
	)

	for i < len(runes) {
		if runes[i] == '_' {
			i++
			upper = true
			continue
		}

		char := runes[i]
		if i == 0 || (runes[i-1] == '_' && upper) {
			char = unicode.ToUpper(char)
		}

		result += string(char)
		i++
		upper = false
	}

	return result
}

func camelCase(val string) string {
	var (
		result = ""
		runes  = []rune(val)
		i      = 0
		upper  = false
	)

	for i < len(runes) {
		if runes[i] == '_' {
			i++
			upper = true
			continue
		}

		char := runes[i]
		if upper && runes[i-1] == '_' {
			char = unicode.ToUpper(char)
		}

		result += string(char)
		i++
		upper = false
	}

	return result
}

func getType(col *column) string {
	var dataType string
	switch col.Type {
	case "TEXT":
		dataType = "string"
	case "INTEGER":
		dataType = "int64"
	case "REAL":
		dataType = "float64"
	case "BLOB":
		return "[]bytes"
	case "NULL":
		return "{}interface"
	}

	if col.IsPrimaryKey {
		return dataType
	}

	if !col.NotNull {
		return fmt.Sprintf("*%s", dataType)
	}

	return dataType
}

func columnNames(cols []*column) []string {
	var items []string
	for _, col := range cols {
		items = append(items, col.Name)
	}

	return items
}

func join(items []string, sep string) string {
	var result string
	for i, item := range items {
		if i > 0 {
			result += sep
		}

		result += item
	}

	return result
}

func mapItems(items []string, val string) []string {
	var result []string
	for range items {
		result = append(result, val)
	}
	return result
}

func filterItems(items []string, excluded ...string) []string {
	var result []string
	for _, item := range items {
		match := false
		for _, ex := range excluded {
			if ex == item {
				match = true
			}
		}

		if !match {
			result = append(result, item)
		}
	}

	return result
}

func backtick() string {
	return "`"
}
