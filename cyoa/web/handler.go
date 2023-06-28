package web

import (
	"encoding/json"
	"github.com/PaulSedra/gophercises/cyoa"
	"html/template"
	"log"
	"net/http"
	"os"
)

// ParseStory parses the JSON file containing the story.
func ParseStory(filename string) (cyoa.Story, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var story cyoa.Story
	err = json.NewDecoder(file).Decode(&story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// HandleStoryArc handles the HTTP request for a specific story arc.
func HandleStoryArc(tmpl *template.Template, story cyoa.Story) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		arc := r.URL.Path[1:] // Remove the leading slash
		storyArc, ok := story[arc]
		if !ok {
			http.NotFound(w, r)
			return
		}

		err := tmpl.Execute(w, storyArc)
		if err != nil {
			log.Println("Failed to execute template:", err)
			http.Error(w, "Oops! Something went wrong.", http.StatusInternalServerError)
		}
	}
}
