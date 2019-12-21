package post

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/google/uuid"
)

var tableName string = "blog_posts"

func FindAll() (Posts, error) {
	svc := createDynamoDBClient()

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		return Posts{}, err
	}

	var posts Posts
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &posts)
	if err != nil {
		return Posts{}, err
	}

	return posts, nil
}

func FindById(id string) (Post, error) {
	svc := createDynamoDBClient()

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return Post{}, err
	}

	// Unmarshall
	item := Post{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return item, fmt.Errorf("Failed to unmarshal Record, %v", err)
	}

	if (Post{}) == item {
		return item, fmt.Errorf("Post not found by id %s", id)
	}

	return item, nil
}

func Save(post Post) (Post, error) {
	dynamodbClient := createDynamoDBClient()

	if post.Id == "" {
		post.Id = uuid.New().String()
		post.CreatedAt = time.Now()
	}

	post.UpdatedAt = time.Now()

	av, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		return Post{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamodbClient.PutItem(input)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	if err != nil {
		panic(err)
	}

	// Create DynamoDB client
	return dynamodb.New(sess)
}
