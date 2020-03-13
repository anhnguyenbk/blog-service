package post // import "github.com/anhnguyenbk/blog-service/internal/post"

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

var tableName string = "blog_posts"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1"))

func FindAll() (Posts, error) {
	// Filter out deleted post
	filt := expression.Name("status").NotEqual(expression.Value("deleted"))

	// Get back the title, year, and rating
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("title"),
		expression.Name("slug"),
		expression.Name("desc"),
		expression.Name("status"),
		expression.Name("categories"),
		expression.Name("createdAt"),
		expression.Name("updatedAt"),
	)

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
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
	result, err := db.GetItem(&dynamodb.GetItemInput{
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

	return item, nil
}

func FindBySlug(slug string) (Post, error) {
	// Create the Expression to fill the input struct with.
	filt := expression.Name("slug").Equal(expression.Value(slug))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		return Post{}, err
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		return Post{}, err
	}

	posts := Posts{}
	fmt.Println(result)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &posts)
	if err != nil {
		return Post{}, err
	}

	if len(posts) == 0 {
		return Post{}, fmt.Errorf("Post not found by slug %s", slug)
	}

	return posts[0], nil
}

func Save(post Post) (Post, error) {
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

	_, err = db.PutItem(input)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func Delete(id string) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String("deleted"),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#ST": aws.String("status"),
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #ST = :s"),
	}

	_, err := db.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}

func FindByCategory(slug string) (Posts, error) {
	// Not supporting: Filtering / Querying by the Contents of a List in DynamoDB
	// https://forums.aws.amazon.com/thread.jspa?messageID=585008&#585008

	posts, err := FindAll()
	if err != nil {
		return nil, err
	}

	result := Posts{}
	for _, post := range posts {
		for _, category := range post.Categories {
			if category.Value == slug {
				result = append(result, post)
			}
		}
	}
	return result, nil
}

func SaveComment(postId string, comment Comment) (Comment, error) {
	comment.Id = uuid.New().String()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	comment.Reply = Comments{}

	dynamoDbItem, err := dynamodbattribute.Marshal(comment)
	if err != nil {
		return Comment{}, err
	}

	var comments []*dynamodb.AttributeValue
	comments = append(comments, dynamoDbItem)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(postId),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":comment": {
				L: comments,
			},
			":empty_list": {
				L: []*dynamodb.AttributeValue{},
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET comments = list_append(if_not_exists(comments, :empty_list), :comment)"),
	}

	_, err = db.UpdateItem(input)
	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func PostGetComments(postId string) (Comments, error) {
	// GetItem with projection expression https://www.dynamodbguide.com/inserting-retrieving-items#get-item
	// Get comments only
	var fields = "comments"
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName:            aws.String(tableName),
		ProjectionExpression: &fields,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(postId),
			},
		},
	})

	if err != nil {
		return Comments{}, err
	}

	//fmt.Println(result)
	// Unmarshall
	item := Post{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return Comments{}, err
	}
	return item.Comments, nil
}

func PostDeleteComment(postId string, commentId string) error {
	post, err := getItemFieldsById(postId, "comments")
	foundCommentIndex := indexOfComment(commentId, post.Comments)

	if foundCommentIndex == -1 {
		return nil
	}

	var updateExpression = "REMOVE comments[" + strconv.Itoa(foundCommentIndex) + "]"
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(postId),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(updateExpression),
	}

	_, err = db.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}

func getItemFieldsById(id string, fields string) (Post, error) {
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName:            aws.String(tableName),
		ProjectionExpression: &fields,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return Post{}, err
	}

	//fmt.Println(result)
	// Unmarshall
	item := Post{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return Post{}, err
	}

	return item, nil
}

func indexOfComment(commentId string, comments []Comment) int {
	if comments == nil {
		return -1
	}

	for k, v := range comments {
		if commentId == v.Id {
			return k
		}
	}
	return -1 //not found.
}
