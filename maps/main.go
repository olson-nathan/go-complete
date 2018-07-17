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
