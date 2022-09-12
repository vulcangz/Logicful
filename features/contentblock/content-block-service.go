package contentblock

import (
	"logicful/models"
	"logicful/service/date"
	"logicful/service/db"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

var instance = db.New()

func Set(block models.ContentBlock) error {

	if block.Id == "" {
		block.Id = uuid.New().String()
	}

	block.CreationDate = date.ISO8601(time.Now())
	block.ChangeDate = date.ISO8601(time.Now())
	block.CreateBy = "maddox"
	block.ChangeBy = "maddox"

	_, err := instance.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Update: &dynamodb.Update{
					TableName: aws.String(db.Clients()),
					Key: map[string]*dynamodb.AttributeValue{
						"name": {
							S: aws.String("maddox"),
						},
					},
					UpdateExpression: aws.String("ADD #20e30 :20e30"),
					ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
						":20e30": {
							SS: aws.StringSlice([]string{block.Id}),
						},
					},
					ExpressionAttributeNames: map[string]*string{
						"#20e30": aws.String("content_blocks"),
					},
				},
			},
			{
				Update: &dynamodb.Update{
					TableName: aws.String(db.ContentBlocks()),
					Key: map[string]*dynamodb.AttributeValue{
						"id": {
							S: aws.String(block.Id),
						},
					},
					UpdateExpression: aws.String("SET #c1e70 = :c1e70, #value = :value, #c1e71 = :c1e71, #c1e73 = if_not_exists(#c1e74,:c1e73), CreateBy = if_not_exists(#CreateBy,:CreateBy), ChangeBy = :ChangeBy"),
					ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
						":c1e70": {
							S: aws.String(block.Name),
						},
						":c1e71": {
							S: aws.String(block.ChangeDate),
						},
						":c1e73": {
							S: aws.String(block.CreationDate),
						},
						":value": {
							S: aws.String(block.Value),
						},
						":ChangeBy": {
							S: aws.String(block.ChangeBy),
						},
						":CreateBy": {
							S: aws.String(block.CreateBy),
						},
					},
					ExpressionAttributeNames: map[string]*string{
						"#c1e70":    aws.String("name"),
						"#c1e71":    aws.String("ChangeDate"),
						"#c1e73":    aws.String("CreationDate"),
						"#c1e74":    aws.String("CreationDate"),
						"#value":    aws.String("value"),
						"#CreateBy": aws.String("CreateBy"),
					},
				},
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func List() ([]models.ContentBlock, error) {
	item, err := instance.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.Clients()),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String("maddox"),
			},
		},
		ProjectionExpression: aws.String("content_blocks"),
	})

	if err != nil {
		return nil, err
	}

	if item == nil || item.Item["content_blocks"] == nil {
		return make([]models.ContentBlock, 0), nil
	}

	ids := item.Item["content_blocks"].SS
	var keys []map[string]*dynamodb.AttributeValue
	for id := range ids {
		k := map[string]*dynamodb.AttributeValue{
			"id": {S: ids[id]},
		}
		keys = append(keys, k)
	}

	items, err := instance.BatchGetItem(&dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			db.ContentBlocks(): {
				Keys: keys,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	results := items.Responses[db.ContentBlocks()]

	var recs []models.ContentBlock

	err = dynamodbattribute.UnmarshalListOfMaps(results, &recs)

	if err != nil {
		return nil, err
	}

	return recs, nil
}
