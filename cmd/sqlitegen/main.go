package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/bencornia/sqlitegen/internal/codegen"
)

func getOrCreateFile(wc *io.WriteCloser) func(string) error {
	return func(s string) error {
		if strings.Compare("", s) == 0 {
			return nil
		}

		f, err := os.Create(s)
		if err != nil {
			return err
		}

		*wc = f
		return nil
	}
}

func main() {
	var wc io.WriteCloser = os.Stdout
	flag.Func("output-file", "Optional output file name (default stdout)", getOrCreateFile(&wc))
	packageName := flag.String("package-name", "model", "Optional package name")
	flag.Parse()

	defer wc.Close()
	codegen.Generate(flag.Args()[0], *packageName, wc)
}
