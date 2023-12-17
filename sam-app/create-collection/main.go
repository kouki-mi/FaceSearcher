package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//postで受け取ったjsonを表示
	fmt.Println(request.Body);
	return events.APIGatewayProxyResponse{
		Body:       "Hello, world!\n",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}