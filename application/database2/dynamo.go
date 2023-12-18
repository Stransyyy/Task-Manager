package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Storage struct {
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

func Open() error {

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-west-2")})
	table := db.Table("tasks")

	return nil

}
