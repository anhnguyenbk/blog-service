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

var posts Posts
var tableName string = "blog_posts"

func FindAll() Posts {
	svc := createDynamoDBClient()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		// ExpressionAttributeNames:  expr.Names(),
		// ExpressionAttributeValues: expr.Values(),
		// FilterExpression:          expr.Filter(),
		// ProjectionExpression:      expr.Projection(),
		TableName: aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &posts)
	if err != nil {
		fmt.Println("Got error unmarshalling:")
		fmt.Println(err.Error())
	}

	return posts
}

func FindById(id string) Post {
	for _, p := range posts {
		if p.Id == id {
			return p
		}
	}
	return Post{}
}

func Create(post Post) Post {
	dynamodbClient := createDynamoDBClient()

	post.Id = uuid.New().String()
	post.Date = time.Now()

	av, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		panic(err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamodbClient.PutItem(input)
	if err != nil {
		panic(err)
	}

	return post
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
