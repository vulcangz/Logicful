package sqs

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"os"
)

func SendMessage(data interface{}, url string) error {
	conf := &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	}
	sess := session.Must(session.NewSession(conf))
	svc := sqs.New(sess)

	serialized, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(uuid.New().String()),
			},
		},
		MessageBody: aws.String(string(serialized)),
		QueueUrl:    aws.String(url),
	})

	if err != nil {
		println(url)
		return err
	}

	return nil
}
