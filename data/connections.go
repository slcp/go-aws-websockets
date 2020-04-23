package data

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ConnectionItemData identifies a connection in the data store
type ConnectionItemData struct {
	ID string
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

func makeConnectionsPK() string {
	return "connectionId"
}

func makeConnectionsSK(now time.Time) string {
	return fmt.Sprintf("%s", string(now.Unix()))
}

func makeConnectionsKey(now time.Time) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"pk": {
			S: aws.String(makeConnectionsPK()),
		},
		"sk": {
			S: aws.String(makeConnectionsSK(now)),
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
