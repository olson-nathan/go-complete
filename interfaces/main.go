package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	// English bot declaration.
	eb := englishBot{}

	// Spanish bot declaration.
	sb := spanishBot{}

	// Print english greeting.
	printGreeting(eb)

	// Print spanish greeting.
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for english implementation.
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for spanish implementation.
	return "Hola!"
}
