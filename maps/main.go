package main

import (
	"fmt"
)

func main() {
	// Create map of [string]string.
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#fbf745",
		"white": "#ffffff",
	}

	// Print out mapssss.
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Color: %v Hex code: %v\n", color, hex)
	}
}

// NOTE: Use a map when you are representing closely related values
// Don't need to know all of the keys at compile time.
// With stucts you have list out property names at compile time.
