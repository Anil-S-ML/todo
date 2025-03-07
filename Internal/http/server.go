package http

import (
	"net/http"

	"todo/internal/manager"
)

type Server struct {
	TodoManager manager.TodoManager
	router      *http.ServeMux
}

func NewServer(m manager.TodoManager) *Server {
	return &Server{
		TodoManager: m,
		router:      http.NewServeMux(),
	}
}

func (s *Server) SetupRoutes() {
	handlers := &Handlers{TodoManager: s.TodoManager}

	// Health check endpoint
	s.router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Todo endpoints
	s.router.HandleFunc("/todos", handlers.handleTodos)
	s.router.HandleFunc("/todos/", handlers.handleTodo)
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}