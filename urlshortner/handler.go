package urlshortner

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"net/http"
)

// a slice of YAML/JSON data
type pathURL struct {
	Path string
	URL  string
}

// YAMLHandler parses the provided YAML and return
// an http.HandlerFunc (which also implements http.Handler)
// that will map any paths to their corresponding URL.
// If the path is not provided in the YAML, then the
// fallback http.Handler will be used.
//
// YAML is expected to be in the format:
//
//	 – path: /example
//		  url: https://www.example.com
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

// JSONHandler parses the provided JSON and return
// an http.HandlerFunc (which also implements http.Handler)
// that will map any paths to their corresponding URL.
// If the path is not provided in the JSON, then the
// fallback http.Handler will be used.
//
// JSON is expected to be in the format:
//
//	{
//	  "path": "/example",
//	  "url": "https://www.example.com/"
//	}
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(jsn)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJSON)
	return MapHandler(pathMap, fallback), nil
}

// MapHandler returns an http.HandlerFunc (which also
// implements http.Handler) that will map any paths to
// their corresponding URL. If the path is not provided
// in the map, then the fallback http.Handler will be used.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

// buildMap creates a map of paths to URL mappings from a slice of pathURL structs
func buildMap(paths []pathURL) map[string]string {
	pathMap := make(map[string]string)
	for _, path := range paths {
		pathMap[path.Path] = path.URL
	}
	return pathMap
}

// parseYAML parses the provided YAML data into a slice of pathURL structs
func parseYAML(yml []byte) ([]pathURL, error) {
	var paths []pathURL
	err := yaml.Unmarshal(yml, &paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// parseJSON parses the provided JSON data into a slice of pathURL structs
func parseJSON(jsn []byte) ([]pathURL, error) {
	var paths []pathURL
	err := json.Unmarshal(jsn, &paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}
