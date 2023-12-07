package dynamo

import (
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamo"
)

var (
client = dynamodb.New(session.New(), aws.NewConfig{
	Region: aws.String("us-east-2"),
})

tableName = "Task-manager"
table = client.Table(tableName)

)
	// Connection is a struct that contains the connection to the DynamoDB database.
func Connection() {