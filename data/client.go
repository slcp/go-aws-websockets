package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DataStore handles daily sales counts for store SKUs.
type DataStore struct {
	Client    *dynamodb.DynamoDB
	TableName *string
}

// CountRecord represents the counts for a SKU
type CountRecord struct {
	SKU   string `json:"sk"`
	Count int64  `json:"count"`
}

// NewDataStore creates a new RoboPharmaIDStore with required fields populated.
func NewDataStore(tableName string, s *session.Session) (store *DataStore) {
	return &DataStore{
		Client:    dynamodb.New(s),
		TableName: aws.String(tableName),
	}
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

// 	var slice []CountRecord
// 	err = dynamodbattribute.UnmarshalListOfMaps(gio.Items, &slice)
// 	if err != nil {
// 		return
// 	}
// 	out = map[string]CountRecord{}
// 	for _, r := range slice {
// 		if ok := strings.HasPrefix(r.SKU, "sku_"); !ok {
// 			return nil, errors.New("unrecognised value for sort key")
// 		}
// 		out[strings.TrimPrefix(r.SKU, "sku_")] = r
// 	}
// 	return
// }
