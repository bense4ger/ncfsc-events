package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bense4ger/ncfsc-events/db"
)

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	evt := db.GetEvent("")

	return events.APIGatewayProxyResponse{
		Body:       evt,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
