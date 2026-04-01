package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	GO PROGRAM 9: HTML RENDERING & TEMPLATES
	Concepts:
	- 'html/template' is Go's template engine for safe and dynamic HTML.
	- Passing data (structs) to templates to render views.
*/

// PageData struct for our template
type PageData struct {
	Title    string
	Author   string
	Language string
	Features []string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Create data to pass to the template
	data := PageData{
		Title:    "Go Programming Language",
		Author:   "Google",
		Language: "Go (Golang)",
		Features: []string{"Concurrency", "Garbage Collection", "Simple Syntax", "Fast Compilation"},
	}
	
	// 2. Define the HTML Template
	tmplStr := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title>
		<style>
			body { font-family: Arial, sans-serif; padding: 2em; background-color: #f4f4f4; }
			h1 { color: #007d9c; }
			ul { color: #333; }
		</style>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<p><strong>Developed by:</strong> {{.Author}}</p>
		<p><strong>Language:</strong> {{.Language}}</p>
		<h3>Key Features:</h3>
		<ul>
			{{range .Features}}
				<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
	</html>
	`
	
	// 3. Parse and execute the template
	tmpl, err := template.New("index").Parse(tmplStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Result is written directly to the ResponseWriter
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Server starting on http://localhost:8081...")
	http.ListenAndServe(":8081", nil)
}
