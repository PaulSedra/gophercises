# Sitemap Builder

The "Sitemap" project provides functionality to build a sitemap by crawling a website and extracting links from it.

## Tools and Concepts Used

To build the Sitemap Builder project, the following tools and concepts were used:

1. URL Parsing: The `net/url` package in Go was used to parse and manipulate URLs. It provides functions to parse URLs from strings, resolve relative URLs, and construct new URLs based on a base URL.
2. HTTP Requests: The `net/http` package in Go was used to send HTTP GET requests to web pages. It allows fetching the HTML content of a webpage, which is then used for link extraction.
3. XML Encoding: The `encoding/xml` package in Go was used to generate XML output for the sitemap. It provides functionality to marshal Go structs into XML format.

## Building the Sitemap

The Sitemap Builder project follows these steps to build the sitemap:

1. URL Enqueueing: The project starts by enqueuing the base URL specified by the user. It maintains a queue to keep track of the URLs to be visited and a visited set to avoid revisiting the same URLs.
2. Crawling and Link Extraction: The project performs a breadth-first search (BFS) traversal of the website's pages up to the specified maximum depth. For each page, it sends an HTTP GET request to retrieve the HTML content. The HTML content is then parsed, and all the links on the page are extracted.
3. URL Filtering: The extracted links are filtered to include only the links that belong to the same domain as the base URL. This ensures that external links are not included in the sitemap.
4. XML Generation: The filtered links are used to build a data structure representing the sitemap. The `encoding/xml` package is used to marshal this data structure into XML format.
5. XML Output: The generated XML sitemap is saved to a file named "sitemap.xml" using the `os` package.