package main

import (
	// "context"
	// "fmt"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
// func serverError(err error) (events.APIGatewayProxyResponse, error) {
//     errorLogger.Println(err.Error())

//     return events.APIGatewayProxyResponse{
//         StatusCode: http.StatusInternalServerError,
//         Body:       http.StatusText(http.StatusInternalServerError),
//     }, nil
// }

// func show(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	posts, err := findAllPost()
// 	if (err != nil) {
// 		return serverError(err)
// 	}

// 	out, err := json.Marshal(posts);
// 	if (err != nil) {
// 		return serverError(err)
// 	}

// 	return events.APIGatewayProxyResponse{
// 		IsBase64Encoded: false,
// 		StatusCode:      200,
// 		Body:            string(out),
// 	}, nil
// }

func main() {
	lambda.Start(Handler)
}
