To build this project, the following tools and concepts were used:

1. Command-Line Flags: The `flag` package in Go was used to define and parse command-line flags. It allowed users to customize the CSV filename, time limit, and shuffle option when running the program.

2. CSV Parsing: The `encoding/csv` package in Go was used to read and parse the CSV file containing the quiz questions and answers. The `ReadAll` function was used to read all records from the CSV file.

3. Randomization: The `math/rand` package in Go was used to shuffle the quiz questions. The `Shuffle` function was used to randomly reorder the records.

4. Time Handling: The `time` package in Go was utilized to handle the time-related functionality of the quiz game. The `NewTimer` function was used to create a timer with a specified time limit. The timer channel (`timer.C`) was used in the `select` statement to check if the time limit had been exceeded.

5. Goroutines: Goroutines were used to implement a non-blocking input mechanism, allowing the program to accept user input while other operations continue to run concurrently.

6. String Manipulation: The `strings` package in Go was used to perform string operations, such as trimming leading and trailing whitespaces from the user's answer and the correct answer.