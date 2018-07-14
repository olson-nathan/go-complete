package main

import "fmt"

// person struct.
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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
	fmt.Printf("%+v\n", alex)

	var mary person

	// Print out mary.
	// firstName and lastName are zero value (string "")
	fmt.Printf("%+v\n", mary)

	// Assign.
	mary.firstName = "Mary"
	mary.lastName = "Magdelene"
	mary.contact.email = "mmagdelene@gmail.com"
	mary.contact.zip = 1234

	// Print.
	fmt.Printf("%+v\n", mary)
}
