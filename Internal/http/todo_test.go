package http

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"todo/internal/auth"
	"todo/internal/manager"
	"todo/internal/todo"

	"github.com/stretchr/testify/assert"
)

func TestTodoHandlers(t *testing.T) {
	tm := manager.NewInMemoryTodoManager()
	authStore := auth.NewInMemoryUserStore()
	authStore.AddUser("testuser", "testpass") // Added back the user creation

	server := httptest.NewServer(NewServer(tm))
	defer server.Close()

	client := server.Client()
	baseURL := server.URL

	t.Run("Create and Retrieve Todo", func(t *testing.T) {
		// Create
		body := `{"title": "Test Todo"}`
		req, _ := http.NewRequest("POST", baseURL+"/todos", strings.NewReader(body))
		req.SetBasicAuth("testuser", "testpass")
		req.Header.Set("Content-Type", "application/json")
		
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		
		var created todo.Todo
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()

		// Retrieve
		req, _ = http.NewRequest("GET", baseURL+"/todos/"+strconv.Itoa(created.ID), nil)
		req.SetBasicAuth("testuser", "testpass")
		
		resp, err = client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		
		var retrieved todo.Todo
		json.NewDecoder(resp.Body).Decode(&retrieved)
		resp.Body.Close()
		
		assert.Equal(t, created.ID, retrieved.ID)
	})

	t.Run("List Todos", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/todos", nil)
		req.SetBasicAuth("testuser", "testpass")
		
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		
		var todos []todo.Todo
		json.NewDecoder(resp.Body).Decode(&todos)
		resp.Body.Close()
		
		assert.GreaterOrEqual(t, len(todos), 1)
	})

	t.Run("Unauthorized Access", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/todos", nil)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Complete Todo", func(t *testing.T) {
		// First create a todo
		body := `{"title": "Complete Test"}`
		req, _ := http.NewRequest("POST", baseURL+"/todos", strings.NewReader(body))
		req.SetBasicAuth("testuser", "testpass")
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req) // Fixed: Added 'err' here
		var created todo.Todo
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()

		// Mark complete
		req, _ = http.NewRequest("PUT", baseURL+"/todos/"+strconv.Itoa(created.ID)+"/complete", nil)
		req.SetBasicAuth("testuser", "testpass")
		resp, err = client.Do(req) // Fixed: Added 'err' here
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Invalid Input Handling", func(t *testing.T) {
		// Test empty title
		body := `{"title": ""}`
		req, _ := http.NewRequest("POST", baseURL+"/todos", strings.NewReader(body))
		req.SetBasicAuth("testuser", "testpass")
		req.Header.Set("Content-Type", "application/json")
		
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestHealthCheck(t *testing.T) {
	tm := manager.NewInMemoryTodoManager()
	server := httptest.NewServer(NewServer(tm))
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "OK", string(body))
}