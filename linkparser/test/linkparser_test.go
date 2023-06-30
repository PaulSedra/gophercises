package test

import (
	"github.com/PaulSedra/gophercises/linkparser/pkg/linkparser"
	"strings"
	"testing"
)

type Link = linkparser.Link

func TestExtractLinks(t *testing.T) {
	testCases := []struct {
		name          string
		html          string
		count         int
		expectedLinks []Link
	}{
		{
			name:  "Single Link",
			html:  `<a href="/test">test</a>`,
			count: 1,
			expectedLinks: []Link{
				{Href: "/test", Text: "test"},
			},
		},
		{
			name:  "Multiple Links",
			html:  `<a href="/test1">test1</a><a href="/test2">test2</a>`,
			count: 2,
			expectedLinks: []Link{
				{Href: "/test1", Text: "test1"},
				{Href: "/test2", Text: "test2"},
			},
		},
		{
			name:  "Nested Links",
			html:  `<a href="/test1">test1<a href="/test2">test2</a></a>`,
			count: 2,
			expectedLinks: []Link{
				{Href: "/test1", Text: "test1"},
				{Href: "/test2", Text: "test2"},
			},
		},
		{
			name:  "Multiple Text Elements",
			html:  `<a href="/test">This is a <strong>strong</strong> test.</a>`,
			count: 1,
			expectedLinks: []Link{
				{Href: "/test", Text: "This is a strong test."},
			},
		},
		{
			name:  "Multiple Text Elements",
			html:  `<a href="/test">test <!-- commented text SHOULD NOT be included! --></a>`,
			count: 1,
			expectedLinks: []Link{
				{Href: "/test", Text: "test"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			links, err := linkparser.ExtractLinks(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("Error extracting links: %v", err)
			}

			if len(links) != tc.count {
				t.Fatalf("Expected %d link, got %d", tc.count, len(links))
			}

			for idx, link := range links {
				expected := tc.expectedLinks[idx]
				if link.Href != tc.expectedLinks[idx].Href {
					t.Errorf("Expected Href %q, got %q", expected.Href, link.Href)
				}
				if link.Text != tc.expectedLinks[idx].Text {
					t.Errorf("Expected Text %q, got %q", expected.Text, link.Text)
				}
			}
		})
	}
}
