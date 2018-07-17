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

	// Update firstName.
	alex.updateFirstName("Tommy")

	// Print.
	alex.print()
}

// Alex is a person struct and not a pointer to person
// Go will automatically translate the variable alex to a pointer.
func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
