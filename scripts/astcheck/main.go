package main

import (
	"log"
	"os"

	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	// "github.com/davecgh/go-spew/spew"
)

func main() {
	// positions are relative to fset
	fset := token.NewFileSet()

	// Read in the original file
	snippet, err := parser.ParseFile(fset, "pkg/engine/info/combat.go", nil, parser.Trace|parser.ParseComments)
	if err != nil {
		log.Panic(err)
	}

	// Print structure to stdout
	ast.Print(fset, snippet)

	// Write to a final file for comparison
	f, err := os.Create("demo.txt")
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()

	err = format.Node(f, fset, snippet)
	if err != nil {
		log.Panic(err)
	}
}
