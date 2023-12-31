package main

import (
	"errors"
	"fmt"

	db "github.com/Stransyyy/Task-Manager/database"
	dynamo "github.com/Stransyyy/Task-Manager/database2"
	task "github.com/Stransyyy/Task-Manager/tsk-mngr"
	"github.com/joho/godotenv"
)

type sliceStorage struct {
	Tasks []*task.Task
}

// type dynamoStorage struct {
// 	db *dynamo.DB
// 	Table dynamo.Table
// }

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

func main() {

	godotenv.Load("run.env")

	// --------------- Slice storage ------------------------------------------------------------------
	//myStorage := sliceStorage{}

	// --------------- MySQL storage ------------------------------------------------------------------

	dbStorage := db.New()

	err := dbStorage.Open()
	if err != nil {
		panic(err)
	}

	// -------------------------------------------------------------------------------------------------

	fmt.Print("\nConnected to database successfully\n\n")

	defer dbStorage.Close()

	// --------------- DynamoDB storage ------------------------------------------------------------------
	dynamoStorage := dynamo.NewDynamo(3)

	err = dynamoStorage.Open("tasks", "us-east-2")
	if err != nil {
		panic(err)
	}

	//-----------------------------------------------------------------------------------------------------
	taskManager := task.Tasks{
		Storage: dynamoStorage,
	}

	taskManager.Run()

}
