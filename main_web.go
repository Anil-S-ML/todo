package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"todo/manager"
)

func main() {
	tm := manager.NewInMemoryTodoManager()

	// Set up routes
	http.HandleFunc("/todos", makeTodosHandler(tm))       // Handles GET/POST for /todos
	http.HandleFunc("/todos/", makeTodoHandler(tm))       // Handles GET/PUT/DELETE for /todos/{id}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	})
	
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

// Handler for /todos endpoint (GET and POST)
func makeTodosHandler(tm *manager.InMemoryTodoManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetTodos(w, r, tm)
		case http.MethodPost:
			handlePostTodo(w, r, tm)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// GET /todos - Retrieve all todos
func handleGetTodos(w http.ResponseWriter, r *http.Request, tm *manager.InMemoryTodoManager) {
	todos := tm.GetAll()
	data, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "Error marshaling todos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// POST /todos - Add a new todo
func handlePostTodo(w http.ResponseWriter, r *http.Request, tm *manager.InMemoryTodoManager) {
	var requestBody struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if requestBody.Title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}
	todo, err := tm.Add(requestBody.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "Error marshaling todo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// Handler for /todos/{id} and /todos/{id}/complete endpoints (GET/PUT/DELETE)
func makeTodoHandler(tm *manager.InMemoryTodoManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		idStr := parts[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			handleGetTodo(w, r, tm, id)
		case http.MethodPut:
			if len(parts) >= 4 && parts[3] == "complete" {
				handlePutComplete(w, r, tm, id)
			} else {
				http.Error(w, "Invalid URL", http.StatusBadRequest)
			}
		case http.MethodDelete:
			handleDeleteTodo(w, r, tm, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// GET /todos/{id} - Retrieve a single todo
func handleGetTodo(w http.ResponseWriter, r *http.Request, tm *manager.InMemoryTodoManager, id int) {
	todo, err := tm.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, "Error marshaling todo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// PUT /todos/{id}/complete - Mark a todo as complete
func handlePutComplete(w http.ResponseWriter, r *http.Request, tm *manager.InMemoryTodoManager, id int) {
	err := tm.MarkComplete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task %d marked as complete", id)
}

// DELETE /todos/{id} - Delete a todo
func handleDeleteTodo(w http.ResponseWriter, r *http.Request, tm *manager.InMemoryTodoManager, id int) {
	err := tm.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task %d deleted", id)
}
