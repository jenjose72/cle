package main

import "fmt"

/*
	GO PROGRAM 1: HELLO WORLD & BASIC OUTPUT
	Concepts:
	- Every runnable file starts with 'package main'.
	- 'import' is used to bring in other packages (like 'fmt' for formatting).
	- 'func main()' is the entry point of every Go application.
*/

func main() {
	// Println adds a newline automatically
	fmt.Println("Welcome to the World of Go!")
	
	// Printf allows for formatted strings
	name := "Gopher"
	fmt.Printf("Hello, %s! Nice to meet you.\n", name)
}
