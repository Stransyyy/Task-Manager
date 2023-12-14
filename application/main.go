package main

import (
	"database/sql"
	"errors"
	"fmt"

	db "github.com/Stransyyy/Task-Manager/mysql"
	task "github.com/Stransyyy/Task-Manager/tsk-mngr"
	"github.com/joho/godotenv"
)

type sliceStorage struct {
	Tasks []*task.Task
}
type dbStorage struct {
	DB *sql.Tx
}

func (s sliceStorage) GetAll() ([]*task.Task, error) {
	return s.Tasks, nil
}

func (s sliceStorage) Store(t *task.Task) error {
	s.Tasks = append(s.Tasks, t)

	return nil
}

func (s sliceStorage) MarkCompleted(id int) error {
	for _, task := range s.Tasks {
		if task.ID == id {
			task.Completed = true
			return nil
		}
	}
	return errors.New("could not find task with that id")
}

func (s sliceStorage) Get(id int) (*task.Task, error) {
	for _, task := range s.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return &task.Task{}, errors.New("could not find task with that id")
}

func (s sliceStorage) Delete(id int) error {
	var newTasks []*task.Task

	for _, task := range s.Tasks {
		if task.ID == id {
			continue
		}

		newTasks = append(newTasks, task)
	}

	return nil
}

func (db dbStorage) GetAll() ([]*task.Task, error) {
	rows, err := db.DB.Query("SELECT id, title, description, due_date, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task

	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}

func (db dbStorage) Store(t *task.Task) error {
	_, err := db.DB.Exec("INSERT INTO tasks (title, description, due_date, completed) VALUES (?, ?, ?, ?)", t.Title, t.Description, t.DueDate, t.Completed)
	return err
}

func (db dbStorage) MarkCompleted(id int) error {
	_, err := db.DB.Exec("UPDATE tasks SET completed = true WHERE id = ?", id)
	return err
}

func (db dbStorage) Delete(id int) error {
	_, err := db.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (db dbStorage) Get(id int) (*task.Task, error) {
	var t task.Task
	err := db.DB.QueryRow("SELECT id, title, description, due_date, completed FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Completed)
	if err != nil {
	}
	return &t, nil
}

func main() {

	godotenv.Load("run.env")

	db, err := db.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Println("")

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	myStorage := sliceStorage{}

	taskManager := task.Tasks{
		Storage: myStorage,
	}

	taskManager.Run()
}
