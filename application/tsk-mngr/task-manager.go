package tm

import (
	"bufio"
	"fmt"
	"os"
)

// This function will create a new task and add it to the database
func new_task() {

}

func edit_task() {

}

func delete_task() {

}

func mark_completed() {

}

func view_tasks() {

}

func task_manager() {

}

// This function will be the main function for the task manager

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()

		fmt.Print("Enter your choice (1-5): ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("You chose Option 1")
			// Add functionality for Option 1 here
		case "2":
			fmt.Println("You chose Option 2")
			// Add functionality for Option 2 here
		case "3":
			fmt.Println("You chose Option 3")
			// Add functionality for Option 3 here
		case "4":
			fmt.Println("You chose Option 4")
			// Add functionality for Option 4 here
		case "5":
			fmt.Println("You chose Option 5")
			// Add functionality for Option 5 here
		case "q":
			fmt.Println("Exiting the scanner. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5 or 'q' to quit.")
		}
	}
}

func printMenu() {
	fmt.Println("Options:")
	fmt.Println("1. Option 1")
	fmt.Println("2. Option 2")
	fmt.Println("3. Option 3")
	fmt.Println("4. Option 4")
	fmt.Println("5. Option 5")
	fmt.Println("q. Quit")
}

/*

	What we need our task manager to do:

	make/delete/edit/mark completed tasks

	Will be visible in the terminal:
	Task Manager
	1. Add Task
	2. View Tasks
	3. Mark Task as Completed
	4. Delete Task
	5. Exit
	Enter your choice (1-5): 1
	Enter the title of the task: spomething
	Enter the description of the task: my description
	Task added successfully!

*/
