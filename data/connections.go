package data

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const (
	ttlAttribute string = "ttl"
	idAttribute  string = "id"
)

// ConnectionItemData identifies a connection in the data store
type ConnectionItemData struct {
	ID string
}

type connectionItemResult struct {
	ID   string `json:"id"`
	Date int    `json:"sk"`
}

// PutConnectionID puts a connectionID into the table.
func (ds Store) PutConnectionID(d ConnectionItemData, now time.Time) (err error) {
	_, err = ds.Client.PutItem(&dynamodb.PutItemInput{
		TableName:    ds.TableName,
		Item:         makeConnectionsItem(d.ID, now),
		ReturnValues: aws.String("NONE"),
	})
	return
}

// GetAllConnectionIDs gets all known connections that were added before now - limit
func (ds Store) GetAllConnectionIDs(now time.Time, limit time.Time) (out []ConnectionItemData, err error) {
	q := expression.Key("pk").Equal(expression.Value(makeConnectionsPK())).
		And(expression.Key("sk").GreaterThan(expression.Value(limit.Unix())))
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

func makeConnectionsPK() *string {
	return aws.String("connectionId")
}

func makeConnectionsSK(now time.Time) *string {
	return timeToUnixTimestampString(now)
}

func makeConnectionsKey(now time.Time) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"pk": {
			S: makeConnectionsPK(),
		},
		"sk": {
			N: makeConnectionsSK(now),
		},
	}
}

func timeToUnixTimestampString(t time.Time) *string {
	return aws.String(fmt.Sprintf("%d", t.Unix()))
}

func makeConnectionsItem(id string, now time.Time) (out map[string]*dynamodb.AttributeValue) {
	out = makeConnectionsKey(now)
	out[idAttribute] = &dynamodb.AttributeValue{
		S: aws.String(id),
	}
	out[ttlAttribute] = &dynamodb.AttributeValue{
		N: timeToUnixTimestampString(now.Add(time.Minute * 20)),
	}
	return
}
