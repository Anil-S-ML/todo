package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
<<<<<<< HEAD:cmd/cli/main.go

<<<<<<< HEAD:cmd/cli/main.go
	"todo/internal/cli"
	"todo/internal/manager"
=======
	"todo/manager" // Relative import to the 'manager' package
	"todo/Utils"   // Relative import to the 'utils' package
>>>>>>> 7beca907d08b58b27200241d60a9c031ec1b458e:main.go
=======
	"todo/manager"
	"todo/Utils"
>>>>>>> 051ff3cc04e000575fa430e7323703d409663f76:main.go
)

func main() {
	fmt.Println("Welcome to the Todo List Application!")
	fmt.Println("You can add multiple tasks. Type 'quit' to exit.")
<<<<<<< HEAD:cmd/cli/main.go

=======
>>>>>>> 7beca907d08b58b27200241d60a9c031ec1b458e:main.go
	todoManager := manager.NewInMemoryTodoManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
<<<<<<< HEAD:cmd/cli/main.go
		task := cli.GetTaskInput(scanner)
		if task == "quit" {
			fmt.Println("Exiting... Here are your tasks:")
<<<<<<< HEAD:cmd/cli/main.go
			todos := todoManager.GetAll()
			cli.PrintTasks(todos)
			cli.MarkTaskComplete(todoManager, scanner)
=======

			todos := manager.GetAll()
			Utils.PrintTasks(todos)
			Utils.MarkTaskComplete(manager, scanner)
>>>>>>> 7beca907d08b58b27200241d60a9c031ec1b458e:main.go
			fmt.Println("Here's your status:")
			updatedTodos := todoManager.GetAll()
			cli.PrintTasks(updatedTodos)

=======

		task := utils.GetTaskInput(scanner)
		if task == "quit" {
			fmt.Println("Exiting... Here are your tasks:")
			todos := todoManager.GetAll()
			utils.PrintTasks(todos)
			utils.MarkTaskComplete(todoManager, scanner)
			fmt.Println("Here's your status:")
			updatedTodos := todoManager.GetAll()
			utils.PrintTasks(updatedTodos)
>>>>>>> 051ff3cc04e000575fa430e7323703d409663f76:main.go

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
