package main

import (
	"flag"
	"fmt"
	"github.com/PaulSedra/gophercises/urlshortner/pkg/db"
	"github.com/PaulSedra/gophercises/urlshortner/pkg/handler"
	"net/http"
	"os"
)

func main() {

	// define YAML/JSON flags
	yamlFile := flag.String("yaml", "urlshortner/redirects.yaml", "path to the YAML file")
	jsonFile := flag.String("json", "urlshortner/redirects.json", "path to the JSON file")
	dbFile := flag.String("db", "urlshortner/redirects.db", "path to the DB file")
	flag.Parse()

	mux := defaultMux()

	// build MapHandler w/ mux as fallback
	pathsToUrls := map[string]string{
		"/gophercises": "https://gophercises.com/",
		"/course":      "https://courses.calhoun.io/courses/cor_gophercises",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// read YAML file
	yamlContent, err := os.ReadFile(*yamlFile)
	if err != nil {
		panic(err)
	}
	// build YAMLHandler w/ MapHandler as fallback
	yamlHandler, err := handler.YAMLHandler(yamlContent, mapHandler)
	if err != nil {
		panic(err)
	}

	// read JSON file
	jsonContent, err := os.ReadFile(*jsonFile)
	if err != nil {
		panic(err)
	}
	// build JSONHandler w/ MapHandler as fallback
	jsonHandler, err := handler.JSONHandler(jsonContent, yamlHandler)
	if err != nil {
		panic(err)
	}

	// open db or create one if it doesn't already exist
	db, err := db.InitializeDB(*dbFile)
	if err != nil {
		panic(err)
	}
	// build DBHandler w/ MapHandler as fallback
	dbHandler := handler.DBHandler(db, jsonHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

// default multiplexer
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultResponse)
	return mux
}

// default response
func defaultResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
