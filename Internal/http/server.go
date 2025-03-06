package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"todo/internal/manager"
)

type Server struct {
	TodoManager manager.TodoManager
}

func NewServer(m manager.TodoManager) *Server {
	return &Server{TodoManager: m}
}

func (s *Server) SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", s.handleTodos)
	mux.HandleFunc("/todos/", s.handleTodo)
	return mux
}

func (s *Server) handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetTodos(w, r)
	case http.MethodPost:
		s.handlePostTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleTodo(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.handleGetTodo(w, r, id)
	case http.MethodPut:
		if len(parts) >= 4 && parts[3] == "complete" {
			s.handlePutComplete(w, r, id)
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}
	case http.MethodDelete:
		s.handleDeleteTodo(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos := s.TodoManager.GetAll()
	data, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, "Error marshaling todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (s *Server) handlePostTodo(w http.ResponseWriter, r *http.Request) {
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

	todo, err := s.TodoManager.Add(requestBody.Title)
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

func (s *Server) handleGetTodo(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := s.TodoManager.Get(id)
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

func (s *Server) handlePutComplete(w http.ResponseWriter, r *http.Request, id int) {
	err := s.TodoManager.MarkComplete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task %d marked as complete", id)
}

func (s *Server) handleDeleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	err := s.TodoManager.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task %d deleted", id)
}