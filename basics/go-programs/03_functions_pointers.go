package main

import "fmt"

/*
	GO PROGRAM 3: FUNCTIONS, MULTIPLE RETURN VALUES & POINTERS
	Concepts:
	- Functions can return multiple results, commonly used for (value, error).
	- Pointers (* and &) allow you to reference memory addresses and modify values directly.
*/

func main() {
	// Function call
	sumResult := add(5, 5)
	fmt.Printf("Add: 5 + 5 = %d\n", sumResult)
	
	// Multiple returns
	mult, div := multiplyAndDivide(10, 2)
	fmt.Printf("Mult: 10 * 2 = %d, Div: 10 / 2 = %d\n", mult, div)
	
	// Pointers
	x := 10
	fmt.Printf("\nOriginal X: %d\n", x)
	
	// Send a pointer (memory address) to the function
	updateByPointer(&x)
	fmt.Printf("After updateByPointer X: %d\n", x)
}

// Simple Function
func add(a int, b int) int {
	return a + b
}

// Multiple Returns Function
func multiplyAndDivide(a int, b int) (int, int) {
	return a * b, a / b
}

// Pointer Function (modifies original variable)
func updateByPointer(val *int) {
	*val = 100 // Overwrites the memory location
}
