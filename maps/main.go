package main

import (
	"fmt"
)

func main() {
	// Multiple ways to create a map in Go.

	// Literal way.
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#fbf745",
	}

	// Variable for map so you can add items later.
	var colors1 map[string]string

	// Another variable declaration that initializes to zero value.
	colors2 := make(map[string]string)

	// Add key/value to map
	// NOTE: this can only be added via bracket notation.
	colors2["white"] = "#ffffff"

	// Delete map key.
	delete(colors2, "white")

	// Print out mapssss.
	fmt.Println("Colors literal map::", colors)
	fmt.Println("Colors map var::", colors1)
	fmt.Println("Colors make::", colors2)

}
