package dynamo

import (
	task "github.com/Stransyyy/Task-Manager/tsk-mngr"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Dynamo struct {
	db      *dynamo.DB
	retries int
}

type Task struct {
	id    string
	title string
}

type dbTaskItem struct {
	PK string `dynamo:"PK"`
	SK string `dynamo:"SK"`
}

var tasks []*Task

func (db Dynamo) GetAll() ([]*Task, error) {

	err := db.db.Table("tasks").Get("PK", "TASK").All(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db Dynamo) Store(t *task.Task) error {

	err := db.db.Table("tasks").Put(&dbTaskItem{
		PK: "TASK",
		SK: t.Title,
	}).Run()
	if err != nil {
		return err
	}

	return nil
}

func (db Dynamo) MarkCompleted(id int) error {
	err := db.db.Table("tasks").Update("PK", "TASK").Range("SK", "TASK").Set("completed", true).Run()
	if err != nil {
		return err
	}

	return nil
}

func (db Dynamo) Delete() error {
	db.db.Table("tasks").Delete("PK", "TASK").Range("SK", "TASK").Run()
	return nil
}

func (db Dynamo) Get(id string) (*task.Task, error) {

	var result dbTaskItem

	err := db.db.Table("tasks").Get("PK", "TASK").Range("SK", dynamo.Operator(id)).One(&result)
	if err != nil {
		return nil, err
	}

	t := &task.Task{
		Title: result.SK,
	}

	return t, err
}

func (db Dynamo) Edit(t *task.Task) error {
	return nil
}

func (db Dynamo) Open() error {

	sess := session.Must(session.NewSession())
	db.db = dynamo.New(sess, &aws.Config{
		Region: aws.String("us-east-1"),
	})

	return nil
}

func NewDynamo() *Dynamo {
	return &Dynamo{}
}
