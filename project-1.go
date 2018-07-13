// Write a quiz program that displays (prompts) a question to the user
// and waits for the user to enter an input (response) to the question
// being asked. Once all questions have been answered the program should
// display the total number of questions the user got correct, over the
// total number of questions asked (i.e You got 5/6 correct).

package main

import "fmt"

var questions = map[string]int{
	"What is 5 + 5?": 10,
	"What is 1 + 1?": 2,
	"What is 8 + 3?": 11,
	"What is 1 + 2?": 3,
	"What is 8 + 6?": 14,
}

func askQuestion(question string) int {
	// Print the question to standard out.
	fmt.Println(question)

	// Scanln scans text read from standard input,
	// storing successive space-separated values into
	// successive arguments as determined by the format.
	// Scanln stops scanning at a newline and after the final
	// item there must be a newline or EOF
	var input int
	fmt.Scanln(&input)

	return input
}

func main() {
	var correct int
	for question, answer := range questions {
		response := askQuestion(question)
		if answer == response {
			correct++
		}
	}
	fmt.Printf("You got %v out of %v correct.\n", correct, len(questions))
}
