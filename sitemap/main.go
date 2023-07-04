package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PaulSedra/gophercises/sitemap/pkg/builder"
)

type Urlset = builder.Urlset

func main() {
	url := flag.String("url", "https://gophercises.com", "starting URL for the sitemap builder")
	maxDepth := flag.Int("depth", 10, "maximum depth to follow links")
	flag.Parse()

	sitemap := builder.BuildSitemap(*url, *maxDepth)
	displayXML(sitemap)
	outputXML(sitemap)
}

func outputXML(sitemap Urlset) {

	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate XML:", err)
		return
	}

	// Open the file for writing, create if it doesn't exist, truncate if it does
	file, err := os.OpenFile("sitemap/sitemap.xml", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// Write the XML to the file
	file.Write([]byte(xml.Header))
	file.Write(output)

	fmt.Println("Sitemap generated and saved to sitemap.xml")
}

func displayXML(sitemap Urlset) {

	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate XML:", err)
		return
	}

	fmt.Println(xml.Header)
	fmt.Println(strings.TrimSpace(string(output)))
}
