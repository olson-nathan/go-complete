package main

import "fmt"

// person struct.
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo will declare a field of contactInfo of type contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

// Main.
func main() {
	// Declare Alex as a person.
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email: "aanderson@gmail.com",
			zip:   89436,
		},
	}

	// Print.
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
