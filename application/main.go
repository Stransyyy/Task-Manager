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

func (db dbStorage) GetAll() (*sql.Tx, error) {
	return db.DB, nil
}

func (s sliceStorage) Store(t *task.Task) error {
	return nil
}

func (s sliceStorage) MarkCompleted(id int) error {

}

func (s sliceStorage) Delete(id int) error {

}

func (s sliceStorage) Get(id int) error {

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
