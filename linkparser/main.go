package main

import (
	"fmt"
	"github.com/PaulSedra/gophercises/linkparser/pkg/linkparser"
	"log"
	"path/filepath"
)

func main() {

	// example files
	path := "linkparser/examples/"
	example := func(n int) string { return fmt.Sprintf("%sex%d.html", path, n) }
	files := []string{example(1), example(2), example(3), example(4)}

	for _, file := range files {
		// extract links from the example files
		links, err := linkparser.ExtractLinksFromFile(file)
		if err != nil {
			log.Fatalf("Failed to extract links from %s: %v", file, err)
		}

		// display the link data in the console
		fmt.Printf("Links extracted from %s:\n", filepath.Base(file))
		for _, link := range links {
			fmt.Printf("Link{\n  Href: %q,\n  Text: %q,\n}\n", link.Href, link.Text)
		}
		fmt.Println()
	}
}
