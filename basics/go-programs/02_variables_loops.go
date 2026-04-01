package main

import "fmt"

/*
	GO PROGRAM 2: VARIABLES, TYPES & THE FOR LOOP
	Concepts:
	- 'var' vs ':=' syntax for declaration.
	- Types: int, string, float64, bool.
	- Loops: Go only has the 'for' keyword, used for standard loops, while-style, and infinite loops.
*/

func main() {
	// Variable declaration with specific type
	var age int = 25
	
	// Shorthand - Go infers the type (only inside functions)
	message := "Your age is:"
	
	// Constants - their values cannot change
	const birthYear = 1999
	
	fmt.Printf("%s %d (Born in %d)\n", message, age, birthYear)
	
	// Standard C-style For Loop
	fmt.Println("\nStandard Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	
	// While-style Loop (using 'for')
	fmt.Println("\nWhile-style Loop:")
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter is %d\n", counter)
		counter++
	}
	
	// Infinite Loop (requires 'break' to stop)
	/* 
	for {
		// break logic
	}
	*/
}
