package main

import (
	"fmt"
)

// Square struct.
type square struct {
	sideLength float64
}

// Triangle struct.
type triangle struct {
	height float64
	base   float64
}

// Shape interface.
type shape interface {
	getArea() float64
}

func main() {
	// Initialize triangle.
	t := triangle{
		height: .4,
		base:   .6,
	}

	// Initialize square.
	s := square{
		sideLength: .5,
	}

	// Print triangle area.
	fmt.Println("Triangle area:", t.getArea())

	// Print square area.
	fmt.Println("Square area:", s.getArea())
}

// Get triangle area.
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.base
}

// Get square area.
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
