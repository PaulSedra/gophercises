package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/PaulSedra/gophercises/cyoa"
	"os"
)

// LoadStory loads the story from a JSON file and returns it as a Story type.
func LoadStory(filename string) (cyoa.Story, error) {

	// open JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decode JSON file into a Story object
	var story cyoa.Story
	err = json.NewDecoder(file).Decode(&story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// PrintArc prints the title and paragraphs of an Arc.
func PrintArc(arc cyoa.Arc) {

	fmt.Println(arc.Title)
	fmt.Println()
	for _, paragraph := range arc.Story {
		fmt.Println(paragraph)
		fmt.Println()
	}
}

// GetUserChoice prompts the user to choose an option from a list and returns the selected choice.
func GetUserChoice(options []cyoa.Option) int {

	// display story options
	fmt.Println("Choose an option:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option.Text)
	}

	// prompt user for choice
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err == nil {
			return choice - 1
		}
		fmt.Println("Invalid choice. Please enter a number.")
	}
}
