package main

import (
	"fmt"
	"net/http"

	"todo/internal/auth"
	myhttp "todo/Internal/http"  // Renamed import to avoid conflict
	"todo/internal/middleware"
	"todo/internal/manager"
)

func main() {
	tm := manager.NewInMemoryTodoManager()
	server := myhttp.NewServer(tm)

	userStore := auth.NewInMemoryUserStore()

	server.SetupRoutes()

	// Apply middleware
	handler := middleware.Logging(server)
	handler = myhttp.AuthMiddleware(userStore)(handler) // Correct reference to http.AuthMiddleware

	http.Handle("/", handler)

	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	http.ListenAndServe(port, nil)
}