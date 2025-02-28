package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *Todo) MarkComplete() {
	t.Completed = true
}

type TodoManager interface {
	Add(title string) (*Todo, error)
	Get(id int) (*Todo, error)
	GetAll() []Todo
	Delete(id int) error
}

type InMemoryTodoManager struct {
	todos  []Todo
	nextID int
}

func NewInMemoryTodoManager() *InMemoryTodoManager {
	return &InMemoryTodoManager{
		todos:  []Todo{},
		nextID: 1,
	}
}

func (tm *InMemoryTodoManager) Add(title string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("task title cannot be empty")
	}
	todo := Todo{
		ID:        tm.nextID,
		Title:     title,
		Completed: false,
	}
	tm.todos = append(tm.todos, todo)
	tm.nextID++
	return &todo, nil
}

func (tm *InMemoryTodoManager) Get(id int) (*Todo, error) {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			return &tm.todos[i], nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (tm *InMemoryTodoManager) GetAll() []Todo {
	return tm.todos
}

func (tm *InMemoryTodoManager) Delete(id int) error {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func main() {
	fmt.Println("Welcome to the Todo List Application!")
	fmt.Println("You can add multiple tasks. Type 'quit' to exit.")
	manager := NewInMemoryTodoManager()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		task := getTaskInput(scanner)
		if task == "quit" {
			fmt.Println("Exiting... Here are your tasks:")
			todos := manager.GetAll()
			printTasks(todos)
			markTaskComplete(manager, scanner)
			fmt.Println("Here's your status:")
			printTasks(todos)

			fmt.Println("Would you like to add more tasks? (yes/no)")
			scanner.Scan()
			if strings.ToLower(scanner.Text()) == "yes" {
				continue
			} else {
				fmt.Println("Exiting the application. Come back again, You need this!")
				break
			}
		}
		_, err := manager.Add(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
		}
	}
}

func getTaskInput(scanner *bufio.Scanner) string {
	fmt.Print("Enter a new task: ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return scanner.Text()
}

func printTasks(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No tasks were added.")
		return
	}

	fmt.Println("\nYour's TO-DO:")
	fmt.Println("| ID   | Task                     | Status          |")
	fmt.Println("|------|--------------------------|-----------------|")

	for _, todo := range todos {
		status := "Not Completed"
		if todo.Completed {
			status = "Completed"
		}
		fmt.Printf("| %-4d | %-24s | %-15s |\n", todo.ID, todo.Title, status)
	}
}

func markTaskComplete(manager TodoManager, scanner *bufio.Scanner) {
	fmt.Print("Enter the IDs of the tasks to mark as completed, separated by commas, or type 'skip' to skip: ")
	scanner.Scan()
	input := scanner.Text()
	if input == "skip" {
		return
	}

	taskIDs := strings.Split(input, ",")
	for _, idStr := range taskIDs {
		idStr = strings.TrimSpace(idStr)
		taskID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", idStr)
			continue
		}

		todo, err := manager.Get(taskID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		todo.MarkComplete()
		fmt.Println("Task marked as completed!")
	}
}