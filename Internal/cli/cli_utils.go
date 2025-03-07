package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings" // Add this import

	"todo/internal/manager"
	"todo/internal/todo"
)

func GetTaskInput(scanner *bufio.Scanner) string {
	fmt.Print("Enter a new task: ")
	scanner.Scan()
	return scanner.Text()
}

func PrintTasks(todos []todo.Todo) {
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

func MarkTaskComplete(manager manager.TodoManager, scanner *bufio.Scanner) {
	fmt.Print("Enter the IDs of the tasks to mark as completed, separated by commas, or type 'skip' to skip: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if input == "skip" {
		return
	}

	var taskIDs []int
	for _, idStr := range strings.Split(input, ",") {
		idStr = strings.TrimSpace(idStr)
		taskID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", idStr)
			continue
		}
		taskIDs = append(taskIDs, taskID)
	}

	for _, taskID := range taskIDs {
		err := manager.MarkComplete(taskID)
		if err != nil {
			fmt.Printf("Error marking task %d: %v\n", taskID, err)
		} else {
			fmt.Printf("Task %d marked as completed!\n", taskID)
		}
	}
}