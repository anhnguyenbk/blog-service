package helper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamoDBSection = CreateDynamoDBSession()

func CreateDynamoDBSession() *dynamodb.DynamoDB {
	return dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1"))
}

func FindByID(tableName string, id string, item interface{}) error {
	err := FindByStringField(tableName, "id", id, item)
	if err != nil {
		return err
	}
	return nil
}

func FindByStringField(tableName string, fieldName string, fieldValue string, item interface{}) error {
	result, err := dynamoDBSection.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			fieldName: {
				S: aws.String(fieldValue),
			},
		},
	})

	if err != nil {
		return err
	}

	// Unmarshall
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return err
	}

	return nil
}
