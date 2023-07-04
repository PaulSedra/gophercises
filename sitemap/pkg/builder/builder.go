package builder

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PaulSedra/gophercises/linkparser/pkg/linkparser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type Loc struct {
	Loc string `xml:"loc"`
}

type Urlset struct {
	Urls  []Loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

// BuildSitemap constructs a sitemap by performing a BFS
// starting from a given URL. It explores the web pages up to a
// specified maximum depth and returns a list of visited URLs.
func BuildSitemap(startURL string, maxDepth int) Urlset {

	visitedURLs := make(map[string]struct{})
	currentLevel := make(map[string]struct{})
	nextLevel := map[string]struct{}{
		startURL: {},
	}
	for depth := 0; depth <= maxDepth; depth++ {
		currentLevel, nextLevel = nextLevel, make(map[string]struct{})
		if len(currentLevel) == 0 {
			break
		}
		for addr := range currentLevel {
			if _, ok := visitedURLs[addr]; ok {
				continue
			}
			visitedURLs[addr] = struct{}{}
			for _, link := range ExtractLinksFromURL(addr) {
				nextLevel[link] = struct{}{}
			}
		}
	}
	pages := make([]string, 0, len(visitedURLs))
	for addr := range visitedURLs {
		pages = append(pages, addr)
	}

	sitemap := Urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		sitemap.Urls = append(sitemap.Urls, Loc{Loc: page})
	}

	return sitemap
}

// ExtractLinksFromURL retrieves the web page at the given
// URL and extracts the links from its content. It returns
// a list of extracted links.
func ExtractLinksFromURL(urlStr string) []string {

	response, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer response.Body.Close()
	requestURL := response.Request.URL
	baseURL := &url.URL{
		Scheme: requestURL.Scheme,
		Host:   requestURL.Host,
	}
	base := baseURL.String()
	return FilterLinks(ExtractHrefs(response.Body, base), HasPrefix(base))
}

// ExtractHrefs extracts the href attributes from the HTML
// content of a web page and converts them to absolute URLs.
// It returns a list of extracted absolute URLs.
func ExtractHrefs(reader io.Reader, base string) []string {

	links, _ := linkparser.ExtractLinks(reader)
	var extractedLinks []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			extractedLinks = append(extractedLinks, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			extractedLinks = append(extractedLinks, link.Href)
		}
	}
	return extractedLinks
}

// FilterLinks filters a list of links based on the given condition.
// It returns a list of links that satisfy the condition.
func FilterLinks(links []string, condition func(string) bool) []string {

	var filteredLinks []string
	for _, link := range links {
		if condition(link) {
			filteredLinks = append(filteredLinks, link)
		}
	}
	return filteredLinks
}

// HasPrefix returns a closure function that checks if a
// link has the specified prefix. The closure function returns
// true if the link has the prefix, and false otherwise.
func HasPrefix(prefix string) func(string) bool {

	return func(link string) bool {
		return strings.HasPrefix(link, prefix)
	}
}
