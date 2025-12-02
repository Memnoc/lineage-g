package main

import (
	"fmt"
	"os"

	"github.com/Memnoc/lineage/internal/parser"
	"github.com/Memnoc/lineage/internal/typst"
)

func main() {
	// Easy way to capture the command line args
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <directory>\n", os.Args[0])
		os.Exit(1)
	}

	p := parser.New()

	if err := p.LoadDirectory(os.Args[1]); err != nil {
		println(os.Args)
		fmt.Fprintf(os.Stderr, "Error loading files: %v\n", err)
		os.Exit(1)
	}

	recipes := p.Process()

	gen := typst.NewGenerator()     // Typst generator
	output := gen.Generate(recipes) // Typst markup

	// TODO: would be nice to have the name of the recipe here
	outputFile := "recipes.typ"
	if err := os.WriteFile(outputFile, []byte(output), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated %s (%d recipes)\n", outputFile, len(recipes))
	fmt.Println("→ Run: typst compile recipes.typ")
}
