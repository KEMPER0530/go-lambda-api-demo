package main

import (
	"log"
	"fmt"
	"os"
	"github.com/joho/godotenv"

	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"go-lambda-api-demo/src/infrastructure"
)

var ginLambda *ginadapter.GinLambda

func init() {
  if os.Getenv("GO_ENV") == "production" {
    //db := infrastructure.NewDB()
    r := infrastructure.NewRouting()
    ginLambda = ginadapter.New(r.Gin)
  }
}

func main() {
	if os.Getenv("GO_ENV") == "production" {
		fmt.Println("Production mode...")
		lambda.Start(Handler)
	}else{
		fmt.Println("Development mode...")
		// 環境変数ファイルの読込
		err := godotenv.Load(fmt.Sprintf("src/infrastructure/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal(err)
		}

		r := infrastructure.NewRouting()
		infrastructure.Run(r)
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}
