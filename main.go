package main

import (
	"context"
	"fmt"
	"user"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var r *gin.Engine
var ginLambda *ginadapter.GinLambda
var db *sqlx.DB

// func init() {
// 	db, _ = sqlx.Open("mysql", "root:password@/db-name")
// }

func init() {
	r = setupRouter()
	ginLambda = ginadapter.New(r)
}

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from AWS Lambda!",
		})
	})

	router.POST("/api/users", user.Create)
	router.GET("/api/users/:id", user.Get)
	router.PUT("/api/users/:id", user.Update)
	router.DELETE("/api/users/:id", user.Delete)

	router.Run(":8000")

	defer db.Close()

	return router
}

func main() {
	fmt.Println("running aws lambda in aws")
	lambda.Start(LambdaHandler)
}
