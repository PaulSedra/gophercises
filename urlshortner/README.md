The "urlshortner" package provides functionality for mapping custom paths to their corresponding URLs using YAML and JSON configuration files, to provide a more concise and user-friendly way of accessing web resources.
A BoltDB solution is also available, allowing the storage of redirect data as key-value pairs in a bucket to be accessed and modified by the application using CRUD.

To build this project, the following tools and concepts were used:

1. Multiplexer: The `net/http` package in Go was used to create an HTTP server and handle incoming requests.
2. Unmarshal: The `encoding/json`, and `pkg.in/yaml.v2` packages in Go were used to parse and unmarshal JSON and YAML data
3. BoldDB: The `github.com/boltdb/bolt` package was used to create a local database as a way to persist data.