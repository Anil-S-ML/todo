package manager

import (
	"sync"
	"errors"
	"fmt"

	"todo/todo"
)

type TodoManager interface {
	Add(title string) (*todo.Todo, error)
	Get(id int) (*todo.Todo, error)
	GetAll() []todo.Todo
	Delete(id int) error
	MarkComplete(id int) error
}

type InMemoryTodoManager struct {
	todos  []todo.Todo
	nextID int
	mu     sync.Mutex
}

func NewInMemoryTodoManager() *InMemoryTodoManager {
	return &InMemoryTodoManager{
		todos:  []todo.Todo{},
		nextID: 1,
	}
}

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

func (tm *InMemoryTodoManager) Get(id int) (*todo.Todo, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.todos {
		if tm.todos[i].ID == id {
			return &tm.todos[i], nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (tm *InMemoryTodoManager) GetAll() []todo.Todo {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	return tm.todos
}

func (tm *InMemoryTodoManager) Delete(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.todos {
		if tm.todos[i].ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tm *InMemoryTodoManager) MarkComplete(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.todos {
		if tm.todos[i].ID == id {
			tm.todos[i].Completed = true
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}