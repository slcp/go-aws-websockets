package data

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// ConnectionItemData identifies a connection in the data store
type ConnectionItemData struct {
	ID string
}

type connectionItemResult struct {
	ID   string `json:"id"`
	Date int    `json:"sk"`
}

// PutConnectionID a connectionID into the table.
func (ds DataStore) PutConnectionID(d ConnectionItemData, now time.Time) (err error) {
	_, err = ds.Client.PutItem(&dynamodb.PutItemInput{
		TableName:    ds.TableName,
		Item:         makeConnectionsItem(d.ID, now),
		ReturnValues: aws.String("NONE"),
	})
	return
}

// TODO: This should get all connections after now
// GetAllConnectionIDs is
func (ds DataStore) GetAllConnectionIDs(now time.Time) (out []ConnectionItemData, err error) {
	q := expression.Key("pk").Equal(expression.Value(makeConnectionsPK()))
	builder, err := expression.NewBuilder().WithKeyCondition(q).Build()
	if err != nil {
		return
	}
	gio, err := ds.Client.Query(&dynamodb.QueryInput{
		ExpressionAttributeValues: builder.Values(),
		ExpressionAttributeNames:  builder.Names(),
		KeyConditionExpression:    builder.KeyCondition(),
		TableName:                 ds.TableName,
	})
	if err != nil {
		return
	}
	var slice []connectionItemResult
	err = dynamodbattribute.UnmarshalListOfMaps(gio.Items, &slice)
	if err != nil {
		return
	}
	out = []ConnectionItemData{}
	for _, i := range slice {
		out = append(out, ConnectionItemData{
			ID: i.ID,
		})
	}
	return
}

// // GetAllDailySalesCounts retrieves daily sales count for a store SKU.
// func (cs DataStore) GetAllDailySalesCounts(date time.Time) (out map[string]CountRecord, err error) {
// 	q := expression.Key("pk").Equal(expression.Value(makePK(date)))
// 	builder, err := expression.NewBuilder().WithKeyCondition(q).Build()
// 	if err != nil {
// 		return
// 	}
// 	gio, err := cs.Client.Query(&dynamodb.QueryInput{
// 		ExpressionAttributeValues: builder.Values(),
// 		ExpressionAttributeNames:  builder.Names(),
// 		KeyConditionExpression:    builder.KeyCondition(),
// 		TableName:                 cs.TableName,
// 	})
// 	if err != nil {
// 		return
// 	}

func makeConnectionsPK() string {
	return "connectionId"
}

func makeConnectionsSK(now time.Time) string {
	return fmt.Sprintf("%d", now.Unix())
}

func makeConnectionsKey(now time.Time) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"pk": {
			S: aws.String(makeConnectionsPK()),
		},
		"sk": {
			N: aws.String(makeConnectionsSK(now)),
		},
	}
}

func makeConnectionsItem(id string, now time.Time) (out map[string]*dynamodb.AttributeValue) {
	out = makeConnectionsKey(now)
	out["id"] = &dynamodb.AttributeValue{
		S: aws.String(id),
	}
	return
}
