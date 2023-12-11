package tm

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Task represents an individual task
type Task struct {
	ID          int       // Unique identifier for the task
	Title       string    // Title of the task
	Description string    // Description of the task
	DueDate     time.Time // Due date of the task
	Completed   bool      // Indicates whether the task is completed
}

// NewTask creates a new Task instance with the provided title, description, and due date
func NewTask(title, description string, dueDate time.Time) *Task {
	return &Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}
}

func newTaskFromUserInput() *Task {
	fmt.Println("Enter the title of the task: ")
	title := readInput()

	fmt.Println("Enter the description of the task: ")
	description := readInput()

	fmt.Println("Enter the due date of the task (YYYY-MM-DD): ")
	dueDateString := readInput()
	dueDate, err := time.Parse("2006-01-02", dueDateString)
	if err != nil {
		fmt.Println("Invalid date format. Task not added.")
		return nil
	}

	return NewTask(title, description, dueDate)
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// This function will create a new task and add it to the database
func new_task(title, description, due_date string) {

	fmt.Println("Enter the title of the task: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Enter the description of the task: ")
	scanner.Scan()
	description = scanner.Text()

	fmt.Println("Enter the due date of the task: ")
	scanner.Scan()
	due_date = scanner.Text()

	fmt.Println("Task added successfully!")

}

func edit_task() {

}

func delete_task() {

}

func mark_completed() {

}

func view_tasks(tasks []Task) {

	if len(tasks) == 0 {
		fmt.Println("No tasks to display")
		return
	}

	fmt.Println("Your tasks: ")
	for _, task := range tasks {

		fmt.Printf("ID: %d\nTitle: %s\nDescription: %s\nDue Date: %s\nCompleted: %t\n", task.ID, task.Title, task.Description, task.DueDate.Format("2006-01-02"), task.Completed)
	}
}

func task_manager() {

}

// This function will be the main function for the task manager

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		displayed()

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

func displayed() {
	fmt.Println("Task manager")
	fmt.Println("1. Add task")
	fmt.Println("2. View tasks")
	fmt.Println("3. Mark task as completed")
	fmt.Println("4. Delete task")
	fmt.Println("5. Edit task")
	fmt.Println("q. Quit / Exit")
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
