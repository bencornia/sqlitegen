package main

import "github.com/bencornia/sqlitegen/pkg/codegen"

func main() {
	codegen.Generate("foo.db", "bar")
}
