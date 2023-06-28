The "Choose Your Own Adventure" (CYOA) project is a web application that aims to recreate the experience of the popular children's book series CYOA. The application allows users to navigate through interactive stories where they make choices that impact the progression of the narrative.

The application implements a web-based interface where each page represents a portion of the story. At the end of each page, the user is presented with a set of options to choose from, determining the next path of the story. If the user reaches the end of a particular story arc, they are notified accordingly.

The project also offers a command-line version of the CYOA application. In this version, the stories are presented in the terminal, and users choose from a set of options by typing in corresponding numbers.

To build the Choose Your Own Adventure (CYOA) project, the following tools and concepts were used:

1. HTML/Template: The `html/template` package in Go was used to create dynamic HTML pages. It provides a powerful templating engine that allows for the seamless integration of Go code within HTML templates.
2. JSON Decoding: The `encoding/json` package in Go was used to decode and parse JSON data. The package facilitated the extraction of story arcs, titles, paragraphs, and options from the JSON file, enabling the dynamic generation of the adventure story.
3. Command-Line Version: The CYOA project also included a command-line version. This version used the terminal as the user interface. This implementation leveraged the capabilities of the Go standard library to handle user input and display story content in the terminal.