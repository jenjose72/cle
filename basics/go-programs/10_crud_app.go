package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

/*
	GO PROGRAM 10: SIMPLE CRUD API (In-Memory)
	Concepts:
	- RESTful API structure.
	- JSON encoding/decoding.
	- Using a Map as a database.
	- Mutex for thread-safety.
*/

// Task model
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// In-Memory "Database"
var (
	tasks  = make(map[int]Task)
	nextID = 1
	mu     sync.Mutex // Ensures only one goroutine modifies the map at a time
)

// Add initial data
func init() {
	tasks[nextID] = Task{ID: nextID, Title: "Learn Go", Completed: false}
	nextID++
}

// Handler for /tasks
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// READ: List all tasks
		mu.Lock()
		taskList := make([]Task, 0, len(tasks))
		for _, t := range tasks {
			taskList = append(taskList, t)
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(taskList)

	case "POST":
		// CREATE: Add a new task
		var nt Task
		if err := json.NewDecoder(r.Body).Decode(&nt); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mu.Lock()
		nt.ID = nextID
		tasks[nextID] = nt
		nextID++
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nt)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler for /tasks/{id}
func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		// READ: Single task
		json.NewEncoder(w).Encode(task)

	case "PUT":
		// UPDATE: Update existing task
		var ut Task
		if err := json.NewDecoder(r.Body).Decode(&ut); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ut.ID = id
		tasks[id] = ut
		json.NewEncoder(w).Encode(ut)

	case "DELETE":
		// DELETE: Remove task
		delete(tasks, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)

	fmt.Println("CRUD Server starting on http://localhost:8082...")
	fmt.Println("Available Endpoints:")
	fmt.Println(" - GET/POST /tasks")
	fmt.Println(" - GET/PUT/DELETE /tasks/{id}")

	http.ListenAndServe(":8082", nil)
}
