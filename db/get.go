package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/bense4ger/ncfsc-events/model"
)

// GetEventByID gets the event with the specified Id
func (c *Client) GetEventByID(id string) (*model.Event, error) {
	keyExp := expression.Key("id").Equal(expression.Value(id))
	proj := createProjection()

	builder := expression.NewBuilder().WithKeyCondition(keyExp).WithProjection(proj)
	exp, err := builder.Build()

	if err != nil {
		return nil, fmt.Errorf("GetEventByID error building expression: %s", err.Error())
	}

	qi := &dynamodb.QueryInput{
		KeyConditionExpression:    exp.KeyCondition(),
		ProjectionExpression:      exp.Projection(),
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		TableName:                 aws.String(c.tableName),
	}

	qo, err := c.client.Query(qi)

	if err != nil {
		return nil, fmt.Errorf("GetEventByID error querying db: %s", err.Error())
	}

	if *qo.Count == 0 {
		return nil, nil
	}

	if *qo.Count > 1 {
		return nil, fmt.Errorf("GetEventByID more than one event returned")
	}

	evt := &model.Event{}

	err = dynamodbattribute.UnmarshalMap(qo.Items[0], evt)
	if err != nil {
		return nil, fmt.Errorf("GetEventById error unmarshalling map: %s", err.Error())
	}

	return evt, nil
}

// GetEvents gets all the events in the db
func (c *Client) GetEvents() (*[]model.Event, error) {
	//So this uses a scan.  Not brilliant with a large data set
	//For this though, we're fine
	si := &dynamodb.ScanInput{
		TableName: aws.String(c.tableName),
	}

	result, err := c.client.Scan(si)
	if err != nil {
		return nil, fmt.Errorf("GetEvents error scanning db: %s", err.Error())
	}

	evts := make([]model.Event, *result.Count)
	for i, e := range result.Items {
		evt := &model.Event{}
		err = dynamodbattribute.UnmarshalMap(e, evt)

		if err != nil {
			//Log out if we error, no need to crash out completely
			log.Printf("GetNextEvent error unmarshalling: %s", err.Error())
			continue
		}

		evts[i] = *evt
	}

	return &evts, nil
}

func createProjection() expression.ProjectionBuilder {
	return expression.NamesList(
		expression.Name("id"),
		expression.Name("dateTime"),
		expression.Name("name"),
		expression.Name("location"),
		expression.Name("info"),
		expression.Name("guests"),
	)
}
