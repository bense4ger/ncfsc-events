package db

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DB is a pointer to a Client
var DB *Client

// Client encapsulates db connection information
type Client struct {
	client    *dynamodb.DynamoDB
	tableName string
}

// CreateClient creates an instance of the Client type
func CreateClient() error {
	cfg := &aws.Config{Region: aws.String("us-east-1")}
	sess := session.Must(session.NewSession(cfg))

	if sess == nil {
		return fmt.Errorf("error creating session for client")
	}

	c := dynamodb.New(sess)

	db := &Client{
		client:    c,
		tableName: os.Getenv("TABLE_NAME"),
	}

	if db.tableName == "" {
		return fmt.Errorf("unable to create DB client - no table name")
	}

	DB = db
	return nil
}
