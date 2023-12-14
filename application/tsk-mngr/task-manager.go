package tm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type TaskStore interface {
	GetAll() ([]*Task, error)
	Store(*Task) error
	MarkCompleted(int) error
	Delete(int) error
	Get(int) (*Task, error)
}

// Task represents an individual task
type Task struct {
	ID          int       // Unique identifier for the task
	Title       string    // Title of the task
	Description string    // Description of the task
	DueDate     time.Time // Due date of the task
	Completed   bool      // Indicates whether the task is completed
}

// NewTask creates a new Task instance with the provided title, description, and due date
func NewTask(title, description string, dueDate time.Time) (*Task, error) {

	if title == "" {
		err := fmt.Errorf("Title cannot be empty. Task not added.")
		return nil, err
	}

	return &Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}, nil
}

// NewTaskFromUserInput creates a new Task instance with the provided title, description, and due date
func newTaskFromUserInput() (*Task, error) {

	fmt.Println("Enter the title of the task: ")
	title := readInput("Title: ")

	if title == "" {
		return nil, fmt.Errorf("Title cannot be empty. Task not added.")
	}

	fmt.Println("Enter the description of the task: ")
	description := readInput("Description: ")

	fmt.Println("Enter the due date of the task (YYYY-MM-DD): ")
	dueDateString := readInput("Due Date: ")
	dueDate, err := time.Parse("2006-01-02", dueDateString)
	if err != nil {
		fmt.Println("Invalid date format. Task not added.")
		return nil, err
	}

	fmt.Printf("Task added successfully!\n\n")

	task, err := NewTask(title, description, dueDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return task, nil
}

type Tasks struct {
	Storage TaskStore
}

// view_tasks will display all the tasks
func (tk Tasks) View() string {
	tasks, err := tk.Storage.GetAll()

	if err != nil {
		return "Error: could not get tasks\n\n"
	}

	if len(tasks) == 0 {
		return "\nNo tasks to display\n\n"
	}
	result := "\nYour tasks: \n"
	for _, task := range tasks {
		result += fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nDue Date: %s\nCompleted: %t\n", task.ID, task.Title, task.Description, task.DueDate.Format("2006-01-02"), task.Completed)
	}
	return result
}

// mark_completed will mark whatever task the user chooses as completed
func (tk Tasks) MarkCompleted(taskID int) error {
	return tk.Storage.MarkCompleted(taskID)
}

// edit_task will edit whatever task the user chooses
func (tk Tasks) Edit_task(taskID int) {

	task, _ := tk.Storage.Get(taskID)

	// for _, task := range tasks {
	// 	if task.ID == taskID {

	fmt.Print("\nEnter the new title of the task\n")
	newTitle := readInput("Title: ")

	fmt.Print("Enter the new description of the task: ")
	newDescription := readInput("Description: ")

	fmt.Print("Enter the new due date of the task (YYYY-MM-DD): ")
	newDueDateString := readInput("Due Date: ")
	newDueDate, err := time.Parse("2006-01-02", newDueDateString)
	if err != nil {
		fmt.Println("Invalid date format. Task not edited.")
		return
	}

	task.Title = newTitle
	task.Description = newDescription
	task.DueDate = newDueDate

	tk.Storage.Store(task)

	fmt.Printf("Task with ID %d edited successfully. \n", taskID)
	return

}

// fmt.Printf("Task with ID %d not found. \n", taskID)

// delete_task will delete whatever task the user chooses
func (tk Tasks) Delete_task(id int) error {
	return tk.Storage.Delete(id)
}

func readTaskID(prompt string) (int, error) {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskID, err := strconv.Atoi((scanner.Text()))
	if err != nil {
		fmt.Println("Error converting string into int: ", err)
		return -1, err
	}
	return taskID, nil
}

// Scans for the user input and returns it
func readInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// This function will be the main function for the task manager
func (tk Tasks) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	var taskIDCounter int

	for {
		displayed()

		fmt.Print("Enter your choice (1-5): ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":

			fmt.Println("Add task")

			task, err := newTaskFromUserInput()

			if err != nil {
				fmt.Println(err)
				continue
			}

			if task != nil {
				task.ID = taskIDCounter
				taskIDCounter++
				tk.Storage.Store(task)
			}

		case "2":

			fmt.Println(tk.View())

		case "3":

			fmt.Println("Mark task as completed")
			fmt.Println("Insert the number of the task you want to mark as completed")
			taskID, err := readTaskID("ID: ")
			if err != nil {
				fmt.Println("Error reading the task: ", err)
				continue
			}
			err = tk.MarkCompleted(taskID)
			if err != nil {
				fmt.Println(err)
				continue
			}

		case "4":
			fmt.Println("Delete task")
			// Add functionality for Option 4 here
		case "5":
			fmt.Println("Edit task")

			fmt.Println("Edit task")
			taskID, err := readTaskID("Enter the ID of the task to edit: ")
			if err != nil {
				fmt.Println("Error editing task:", err)
				return
			}

			tk.Edit_task(taskID)

		case "q":
			fmt.Println("Exiting task manager...")
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
