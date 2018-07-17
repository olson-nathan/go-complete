package main

import (
	"fmt"
)

// NOTE: an interface is not a concrete type.
// An interface is an interface type.
// Interfaces are implicit in Go, not explicit.
type bot interface {
	getGreeting() string
}

// NOTE: a map, struct, string is a concrete type.
// This is because we can create a value out of it.
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
