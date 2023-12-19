package dynamo

import (
	"fmt"

	task "github.com/Stransyyy/Task-Manager/tsk-mngr"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const (
	pk         = "PK"
	sk         = "SK"
	allTasksPK = "AllTasks"
)

type Dynamo struct {
	db      *dynamo.DB
	table   dynamo.Table
	retries int
}

type dbTaskItem struct {
	PK string `dynamo:"PK"`
	SK string `dynamo:"SK"`
	task.Task
}

func (db *Dynamo) GetAll() ([]*task.Task, error) {

	var tasks []*task.Task

	err := db.table.Get(pk, allTasksPK).All(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func taskKey(id int) (string, string) {
	return fmt.Sprintf("%d", id), "TASK"
}

func (db *Dynamo) Store(t *task.Task) error {
	taskPK, taskSK := taskKey(t.ID)

	tx := db.db.WriteTx()

	tx.Put(db.table.Put(&dbTaskItem{
		PK:   taskPK,
		SK:   taskSK,
		Task: *t,
	}))

	tx.Put(db.table.Put(&dbTaskItem{
		PK:   allTasksPK,
		SK:   taskPK,
		Task: *t,
	}))

	err := tx.Run()

	if err != nil {
		return err
	}

	return nil
}

func (db *Dynamo) MarkCompleted(id int) error {
	taskPK, taskSK := taskKey(id)

	tx := db.db.WriteTx()

	tx.Update(db.table.Update(pk, taskPK).Range(sk, taskSK).Set("Completed", true))

	tx.Update(db.table.Update(pk, allTasksPK).Range(sk, taskPK).Set("Completed", true))

	err := tx.Run()

	if err != nil {
		return err
	}

	return nil
}

func (db *Dynamo) Delete(id int) error {
	taskPK, taskSK := taskKey(id)

	tx := db.db.WriteTx()

	tx.Delete(db.table.Delete(pk, taskPK).Range(sk, taskSK))

	tx.Delete(db.table.Delete(pk, allTasksPK).Range(sk, taskPK))

	err := tx.Run()

	if err != nil {
		return err
	}

	return nil
}

func (db *Dynamo) Get(id int) (*task.Task, error) {

	taskPK, taskSK := taskKey(id)

	var result dbTaskItem

	err := db.table.Get(pk, taskPK).Range(sk, dynamo.Equal, taskSK).One(&result)
	if err != nil {
		return nil, err
	}

	return &result.Task, nil
}

func (db Dynamo) Edit(t *task.Task) error {
	return db.Store(t)
}

func (db *Dynamo) Open(table, region string) error {

	sess := session.Must(session.NewSession())
	db.db = dynamo.New(sess, &aws.Config{
		Region: aws.String(region),
	})

	db.table = db.db.Table(table)

	return nil
}

func (db *Dynamo) Close() error {

	db.db = nil

	db.table = dynamo.Table{}

	return nil
}

func NewDynamo(retries int) *Dynamo {
	return &Dynamo{
		retries: retries,
	}
}
