package manager

import (
	"errors"
	"fmt"
	"sync"

<<<<<<< HEAD:Internal/manager/todo_manager.go
	"todo/internal/todo"
=======
	// Relative import to the 'todo' package
	"todo/todo"
>>>>>>> 051ff3cc04e000575fa430e7323703d409663f76:manager/todo_manager.go
)

// TodoManager interface defines methods for managing todos.
type TodoManager interface {
	Add(title string) (*todo.Todo, error)
	Get(id int) (*todo.Todo, error)
	GetAll() []todo.Todo
	Delete(id int) error
    MarkComplete(id int) error

}


// InMemoryTodoManager implements the TodoManager interface with an in-memory list.
type InMemoryTodoManager struct {
	todos  []todo.Todo
	nextID int
	mu     sync.Mutex
}

// NewInMemoryTodoManager creates a new instance of InMemoryTodoManager.
func NewInMemoryTodoManager() *InMemoryTodoManager {
	return &InMemoryTodoManager{
		todos:  make([]todo.Todo, 0),
		nextID: 1,
	}
}

// Add adds a new todo item to the list.
func (tm *InMemoryTodoManager) Add(title string) (*todo.Todo, error) {
    tm.mu.Lock()
    defer tm.mu.Unlock()
	if title == "" {
		return nil, errors.New("task title cannot be empty")
	}
	todo := todo.Todo{
		ID:        tm.nextID,
		Title:     title,
		Completed: false,
	}
	tm.todos = append(tm.todos, todo)
	tm.nextID++
	return &todo, nil
}

// Get retrieves a todo by its ID.
func (tm *InMemoryTodoManager) Get(id int) (*todo.Todo, error) {
<<<<<<< HEAD:Internal/manager/todo_manager.go
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for _, t := range tm.todos {
		if t.ID == id {
			return &t, nil
=======
    tm.mu.Lock()
    defer tm.mu.Unlock()
    	for i := range tm.todos {
		if tm.todos[i].ID == id {
			return &tm.todos[i], nil
>>>>>>> 051ff3cc04e000575fa430e7323703d409663f76:manager/todo_manager.go
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

// GetAll retrieves all todos.
func (tm *InMemoryTodoManager) GetAll() []todo.Todo {
    tm.mu.Lock()
    defer tm.mu.Unlock()
	return tm.todos
}

// Delete deletes a todo by its ID.
func (tm *InMemoryTodoManager) Delete(id int) error {
<<<<<<< HEAD:Internal/manager/todo_manager.go
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i, t := range tm.todos {
		if t.ID == id {
=======
    tm.mu.Lock()
    defer tm.mu.Unlock()
	for i := range tm.todos {
		if tm.todos[i].ID == id {
>>>>>>> 051ff3cc04e000575fa430e7323703d409663f76:manager/todo_manager.go
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
func (tm *InMemoryTodoManager) MarkComplete(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i, t := range tm.todos {
		if t.ID == id {
			tm.todos[i].Completed = true
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}