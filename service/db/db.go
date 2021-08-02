package db

import (
	"cloud.google.com/go/firestore"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"golang.org/x/net/context"
	"log"
	"os"
)

var client *firestore.Client

func Submissions() string {
	return os.Getenv("FORM_SUBMISSIONS_TABLE")
}

func Forms() string {
	return os.Getenv("FORMS_TABLE")
}

func Data() string {
	return os.Getenv("DATA_TABLE")
}

func Clients() string {
	return os.Getenv("CLIENTS_TABLE")
}

func OptionSets() string {
	return os.Getenv("OPTION_SETS_TABLE")
}

func DistributedLocks() string {
	return os.Getenv("SUBMISSION_QUEUE_LOCKS_TABLE")
}

func ContentBlocks() string {
	return os.Getenv("CONTENT_BLOCKS_TABLE")
}

func New() *dynamodb.DynamoDB {
	conf := aws.Config{
		Region: aws.String(os.Getenv("DYNAMO_REGION")),
	}
	local := os.Getenv("DYNAMO_IS_LOCAL")
	if local == "true" {
		conf.Endpoint = aws.String(os.Getenv("DYNAMO_ENDPOINT"))
	}
	sess := session.Must(session.NewSession())
	// Create a DynamoDB team with additional configuration
	return dynamodb.New(sess, &conf)
}

func Instance() *firestore.Client {
	if client != nil {
		return client
	}
	projectID := "logicful-290101"
	var err error
	client, err = firestore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}
