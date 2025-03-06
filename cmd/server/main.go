package main

import (
	"fmt"
	"log"
	
	stdhttp "net/http"

	"todo/internal/http"
	"todo/internal/manager"
)

func main() {
	tm := manager.NewInMemoryTodoManager()
	server := http.NewServer(tm)
	
	stdhttp.Handle("/", server.SetupRoutes())
	
	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	log.Fatal(stdhttp.ListenAndServe(port, nil))
}