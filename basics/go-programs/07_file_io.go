package main

import (
	"fmt"
	"os"
	"bufio"
)

/*
	GO PROGRAM 7: FILE I/O (READING & WRITING)
	Concepts:
	- 'os' package for file operations.
	- 'bufio' for buffered reading/writing.
	- 'defer' is used to ensure a cleanup operation (like closing a file) happens at the end of the function.
*/

func main() {
	filename := "testfile.txt"
	content := "Line 1: Go is awesome!\nLine 2: File handling is easy too."
	
	// WRITE FILE
	fmt.Printf("Writing to: %s\n", filename)
	
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	
	// READ FILE
	fmt.Println("Reading from file:")
	
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // ALWAYS Close files when finished!
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Line found: %s\n", scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error during scan: %v\n", err)
	}
}
