package data_test

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateTestTable() (sess *session.Session, tableName string, drop func() error, err error) {
	config := aws.NewConfig().
		WithEndpoint("http://localhost:8000").
		WithRegion("eu-west-2").
		WithCredentials(credentials.NewStaticCredentials("dummy", "dummy", ""))
	sess, err = session.NewSessionWithOptions(session.Options{
		Config: *config,
	})
	if err != nil {
		return
	}
	tableName = fmt.Sprintf("test_connections_%d", time.Now().UnixNano())
	client := dynamodb.New(sess)
	_, err = client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("pk"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("sk"), AttributeType: aws.String("S")},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("pk"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("sk"), KeyType: aws.String("RANGE")},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
	})
	if err != nil {
		return
	}
	drop = func() error {
		return dropTable(client, tableName)
	}
	return
}

func dropTable(client *dynamodb.DynamoDB, name string) error {
	_, err := client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(name),
	})
	return err
}
