//go:generate go run generate.go

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/simimpact/srsim/gen"
)

func main() {
	if !gen.IsDMAvailable() {
		return
	}

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	chars := gen.GetCharacters()

	for _, e := range entries {
		if !e.IsDir() || e.Name() == "dummy" {
			continue
		}

		out, err := os.Create(filepath.Join(e.Name(), "data.go"))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		gen.GenerateCharPromotions(out, chars[e.Name()])
	}
}
