// Package linkparser provides functionality to parse HTML files and extract links.
package linkparser

import (
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

// Link represents a parsed link containing the href and text.
type Link struct {
	Href string
	Text string
}

// ExtractLinksFromFile extracts links from the specified HTML file.
// It returns a slice of Link objects representing the extracted links.
// If an error occurs while opening or reading the file, it returns the error.
func ExtractLinksFromFile(filename string) ([]Link, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ExtractLinks(file)
}

// ExtractLinks extracts links from the provided io.Reader.
// It parses the HTML content and traverses the HTML nodes to extract links.
// It returns a slice of Link objects representing the extracted links.
// If an error occurs while parsing the HTML, it returns the error.
func ExtractLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []Link
	var dfs func(*html.Node)
	dfs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			link := extractLink(n)
			if link.Href != "" && link.Text != "" {
				link.Text = strings.TrimSpace(link.Text)
				links = append(links, link)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			dfs(c)
		}
	}
	dfs(doc)

	return links, nil
}

// extractLink extracts the href and text from the provided HTML node.
// It returns a Link object representing the extracted link.
func extractLink(n *html.Node) Link {
	var link Link

	var getText func(*html.Node)
	getText = func(n *html.Node) {
		if n.Type == html.TextNode {
			link.Text += n.Data
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				getText(c)
			}
		}
	}

	getText(n)

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}

	return link
}
