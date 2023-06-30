The "HTML Link Parser" package provides functionality to parse HTML files and extract links from them.

To build the HTML Link Parser project, the following tools and concepts were used:

1. File Paths: The `path/filepath` package in Go (specifically the `filepath.Base` function) was used to extract the base name (file name) from the full file path.
2. HTML Parsing: The `golang.org/x/net/html` package was used to parse the HTML content and traverse the HTML nodes to extract links.
3. Testing: The `testing` in Go was used to write unit tests to validate the functionality of the link parser. The `go test` command is used to execute these tests and ensure the correctness of the code.