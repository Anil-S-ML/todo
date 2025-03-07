package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"todo/internal/manager"
)

type Handlers struct {
	TodoManager manager.TodoManager
}

func (h *Handlers) handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetTodos(w, r)
	case http.MethodPost:
		h.handlePostTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) handleTodo(w http.ResponseWriter, r *http.Request) {
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
		h.handleGetTodo(w, r, id)
	case http.MethodPut:
		if len(parts) >= 4 && parts[3] == "complete" {
			h.handlePutComplete(w, r, id)
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}
	case http.MethodDelete:
		h.handleDeleteTodo(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.TodoManager.GetAll()
	respondJSON(w, todos)
}

func (h *Handlers) handlePostTodo(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	title := strings.TrimSpace(requestBody.Title)
	if title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}
	if len(title) > 100 {
		http.Error(w, "Title too long (max 100 characters)", http.StatusBadRequest)
		return
	}

	todo, err := h.TodoManager.Add(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, todo, http.StatusCreated)
}

func (h *Handlers) handleGetTodo(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := h.TodoManager.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJSON(w, todo)
}

func (h *Handlers) handlePutComplete(w http.ResponseWriter, r *http.Request, id int) {
	err := h.TodoManager.MarkComplete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) handleDeleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	err := h.TodoManager.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func respondJSON(w http.ResponseWriter, data interface{}, status ...int) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}