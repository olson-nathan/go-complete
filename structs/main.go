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

	var mary person

	// Print out mary.
	// firstName and lastName are zero value (string "")
	fmt.Println(mary)

	// Assign.
	mary.firstName = "Mary"
	mary.lastName = "Magdelene"

	// Print.
	fmt.Println(mary)
}
