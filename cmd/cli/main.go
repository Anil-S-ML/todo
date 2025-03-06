package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo/internal/cli"
	"todo/internal/manager"
)

func main() {
	fmt.Println("Welcome to the Todo List Application!")
	fmt.Println("You can add multiple tasks. Type 'quit' to exit.")

	todoManager := manager.NewInMemoryTodoManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		task := cli.GetTaskInput(scanner)
		if task == "quit" {
			fmt.Println("Exiting... Here are your tasks:")
			todos := todoManager.GetAll()
			cli.PrintTasks(todos)
			cli.MarkTaskComplete(todoManager, scanner)
			fmt.Println("Here's your status:")
			updatedTodos := todoManager.GetAll()
			cli.PrintTasks(updatedTodos)

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