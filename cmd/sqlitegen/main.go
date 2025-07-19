package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bencornia/sqlitegen/pkg/codegen"
)

func main() {
	outFile := flag.String("output", "", "Optional output file name")
	packageName := flag.String("package-name", "model", "Optional package name")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage %s [-output <outputfile>] <inputfile>\n", os.Args[0])
		os.Exit(1)
	}

	var writer io.Writer = os.Stdout
	if *outFile != "" {
		file, err := os.Create(*outFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
			os.Exit(1)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "An unknown error occurred while closing the output file: %v\n", err)
				os.Exit(1)
			}
		}(file)

		writer = file
	}

	codegen.Generate(args[0], *packageName, writer)
}
