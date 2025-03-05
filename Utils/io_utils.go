package Utils

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "todo/todo"     // Import todo package to use Todo struct
    "todo/manager"  // Import manager package to interact with TodoManager
)
func GetTaskInput(scanner *bufio.Scanner) string {
    fmt.Print("Enter a new task: ")
    scanner.Scan()
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input:", err)
    }
    return scanner.Text()
}

func PrintTasks(todos []todo.Todo) {  // <-- Change to todo.Todo
    if len(todos) == 0 {
        fmt.Println("No tasks were added.")
        return
    }

    fmt.Println("\nYour TO-DO list:")
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
