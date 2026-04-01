package main

import "fmt"

/*
	GO PROGRAM 4: STRUCTS, METHODS & INTERFACES
	Concepts:
	- Structs: Grouped data fields.
	- Methods: Functions attached to structs.
	- Interfaces: Define a behavior (set of methods) that any type can satisfy.
*/

// Shape Interface defines anything that can calculate its area
type Shape interface {
	Area() float64
}

// Rectangle Struct
type Rectangle struct {
	Width, Height float64
}

// Circle Struct
type Circle struct {
	Radius float64
}

// Area Method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area Method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Function that uses the Interface - it doesn't care if it's a rectangle or a circle!
func printArea(s Shape) {
	fmt.Printf("Area of the shape: %.2f\n", s.Area())
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}
	
	fmt.Println("Working with Shapes (Structs & Interfaces):")
	
	// Both Rectangle and Circle "implement" the Shape interface because they have an Area() method.
	printArea(r)
	printArea(c)
}
