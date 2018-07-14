package main

import "fmt"

// person struct.
type person struct {
	firstName string
	lastName  string
}

// Main.
func main() {
	// Declare Alex as a person.
	alex := person{firstName: "Alex", lastName: "Anderson"}

	// Print.
	fmt.Println(alex)
}
