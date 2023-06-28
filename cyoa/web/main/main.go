package main

import (
	"github.com/PaulSedra/gophercises/cyoa/web"
	"html/template"
	"log"
	"net/http"
)

func main() {

	story, err := web.ParseStory("cyoa/gopher.json")
	if err != nil {
		log.Fatal("Failed to parse story:", err)
	}

	tmpl := template.Must(template.ParseFiles("cyoa/web/template.html"))

	http.HandleFunc("/", web.HandleStoryArc(tmpl, story))

	log.Println("Starting server on http://localhost:8080/intro")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
