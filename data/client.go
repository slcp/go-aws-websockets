package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Store handles daily sales counts for store SKUs.
type Store struct {
	Client    *dynamodb.DynamoDB
	TableName *string
}

// NewStore creates a new Store ready to access a table.
func NewStore(tableName string, s *session.Session) (store *Store) {
	return &Store{
		Client:    dynamodb.New(s),
		TableName: aws.String(tableName),
	}
}
