package main

import (
	"errors"
	"fmt"
)

/*
	GO PROGRAM 5: ERROR HANDLING
	Concepts:
	- Go uses a "return error" approach, not try-catch.
	- Usually, functions return (Result, error).
	- If the second return is not nil, an error occurred.
*/

func divide(a, b float64) (float64, error) {
	// Custom error if denominator is zero
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Success logic
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}
	
	// Error logic
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("Something went wrong (as expected): %v\n", err2)
	} else {
		fmt.Printf("Result is: %.2f\n", result2)
	}
}
