package main

import (
	"fmt"
	"net/http"
)

/*
	GO PROGRAM 8: BASIC HTTP SERVER
	Concepts:
	- 'net/http' is Go's powerful standard library for web servers.
	- 'http.HandleFunc()' registers a function to handle a specific URL path.
	- 'http.ListenAndServe()' starts the server on a given port.
*/

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Send a plain text response
	fmt.Fprintf(w, "Welcome to the Go Home Page!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Extract a query parameter (e.g., /hello?name=Gopher)
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	
	// Start the server
	fmt.Println("Server starting on port 8080...")
	fmt.Println("Try visiting http://localhost:8080 or http://localhost:8080/hello?name=YourName")
	
	// Listen and serve on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
