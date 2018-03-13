package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bense4ger/ncfsc-events/db"
)

func getEvents() (string, error) {
	evts, err := db.DB.GetEvents()
	if err != nil {
		return "", fmt.Errorf("error getting events: %s", err.Error())
	}

	b, err := json.Marshal(evts)
	if err != nil {
		return "", fmt.Errorf("error marshalling events: %s", err.Error())
	}

	return string(b), nil
}

func getEvent(id string) (string, error) {
	evt, err := db.DB.GetEventByID(id)
	if err != nil {
		return "", fmt.Errorf("error getting event: %s", err.Error())
	}

	b, err := json.Marshal(evt)
	if err != nil {
		return "", fmt.Errorf("error marshalling event: %s", err.Error())
	}

	return string(b), nil
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := db.CreateClient()
	if err != nil {
		log.Printf("handleRequest error creating client: %s", err.Error())

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, fmt.Errorf(err.Error())
	}

	id := req.QueryStringParameters["id"]

	var (
		response string
		respErr  error
	)

	if id == "" {
		response, respErr = getEvents()
	} else {
		response, respErr = getEvent(id)
	}

	if respErr != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, respErr
	}

	var statusCode int
	if response == "{}" {
		statusCode = 201
	} else {
		statusCode = 200
	}

	return events.APIGatewayProxyResponse{
		Body:       response,
		StatusCode: statusCode,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
