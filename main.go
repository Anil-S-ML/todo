package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo/manager" // Relative import to the 'manager' package
	"todo/Utils"   // Relative import to the 'utils' package
)

func main() {
	fmt.Println("Welcome to the Todo List Application!")
	fmt.Println("You can add multiple tasks. Type 'quit' to exit.")
	todoManager := manager.NewInMemoryTodoManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		task := Utils.GetTaskInput(scanner)
		if task == "quit" {
			fmt.Println("Exiting... Here are your tasks:")

			todos := manager.GetAll()
			Utils.PrintTasks(todos)
			Utils.MarkTaskComplete(manager, scanner)
			fmt.Println("Here's your status:")
			Utils.PrintTasks(todos)


			fmt.Println("Would you like to add more tasks? (yes/no)")
			scanner.Scan()
			if strings.ToLower(scanner.Text()) == "yes" {
				continue
			} else {
				fmt.Println("Exiting the application. Come back again, You need this!")
				break
			}
		}

		_, err := todoManager.Add(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
		}
	}
}