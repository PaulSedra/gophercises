package main

import (
	"fmt"
	"github.com/PaulSedra/gophercises/cyoa/cmd"
	"log"
)

func main() {

	story, err := cmd.LoadStory("cyoa/gopher.json")
	if err != nil {
		log.Fatal("Failed to load story:", err)
	}

	currentArc := "intro"

	for len(story[currentArc].Options) > 0 {
		cmd.PrintArc(story[currentArc])
		optionIndex := cmd.GetUserChoice(story[currentArc].Options)
		if optionIndex < 0 || optionIndex >= len(story[currentArc].Options) {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
		currentArc = story[currentArc].Options[optionIndex].Arc
	}

	cmd.PrintArc(story[currentArc])
	fmt.Println("The end.")
}
