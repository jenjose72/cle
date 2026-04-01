# Go Programming Examples

This directory contains 10 different Go programs designed to teach you the core concepts of the language, from basic syntax to building a web server and a CRUD API.

## How to Run the Programs
Ensure you have [Go installed](https://go.dev/doc/install) on your system.

To run any of these programs, open your terminal in this directory and use:
```bash
go run <filename>.go
```

## Program Overview

1. **01_hello_world.go**: The classic entry point to Go, covering basic output and packages.
2. **02_variables_loops.go**: Learn about variable types, constants, and the flexible 'for' loop.
3. **03_functions_pointers.go**: Functions, multiple return values, and how to use pointers.
4. **04_structs_methods_interfaces.go**: Object-oriented concepts in Go using structs and interfaces.
5. **05_error_handling.go**: The Go idiomatic way of handling errors without try-catch.
6. **06_concurrency.go**: Mastering Go's power with goroutines, channels, and wait groups.
7. **07_file_io.go**: Reading from and writing to the local filesystem.
8. **08_basic_server.go**: A simple web server using the `net/http` standard library.
9. **09_html_rendering.go**: Serving dynamic HTML using Go's `html/template` engine.
10. **10_crud_app.go**: A complete in-memory CRUD API for a "Task" resource.

### Note on Web Servers
For programs 8, 9, and 10, once you run them, they will start a server. You can visit them in your browser at:
- **08_basic_server.go**: `http://localhost:8080`
- **09_html_rendering.go**: `http://localhost:8081`
- **10_crud_app.go**: `http://localhost:8082`

Happy Coding!
