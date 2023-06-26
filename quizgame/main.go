package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	// define command-line flags
	csvFilename := flag.String("csv", "problems.csv", "the CSV file containing the quiz")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", true, "whether to shuffle the quiz questions")
	flag.Parse()

	// read CSV
	records := readCSV(csvFilename, shuffle)

	// prompt user to start quiz
	fmt.Println("Press Enter to start the quiz.")
	fmt.Scanln()

	// create timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// run quiz
	correctCount := quiz(records, timer)

	// quiz completed display results
	printQuizResults(correctCount, len(records))
}

func readCSV(csvFilename *string, shuffle *bool) [][]string {

	// open CSV file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	defer file.Close()

	// read CSV records
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	// shuffle quiz questions
	if *shuffle {
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	return records
}

func quiz(records [][]string, timer *time.Timer) int {

	correctCount := 0

	// iterate over the quiz questions
	for i, record := range records {

		question := record[0]
		answer := strings.TrimSpace(record[1])

		// print question
		fmt.Printf("Problem #%d: %s = ", i+1, question)

		// create a channel to receive the user's answer
		answerCh := make(chan string)

		// Goroutine to read the user's answer
		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			answerCh <- strings.TrimSpace(userAnswer)
		}()

		select {
		case <-timer.C: // timer channel
			// timer expired
			fmt.Println("\nTime's up!")
			return correctCount
		case userAnswer := <-answerCh: // answer channel
			// user provided answer
			if userAnswer == answer {
				correctCount++
			} // answer is correct
		}
	}

	return correctCount
}

func printQuizResults(correctCount int, totalQuestions int) {
	fmt.Printf("You scored %d out of %d.\n", correctCount, totalQuestions)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
