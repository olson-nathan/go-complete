package main

import "fmt"

func main() {
	// Create int slice.
	is := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Loop over int slice.
	for _, v := range is {
		// Evaluate if value modulo 2 has a remainder of zero.
		if v%2 == 0 {
			fmt.Printf("%v is even\n", v)
		} else {
			// Otherwise odd.
			fmt.Printf("%v is odd\n", v)
		}
	}
}
