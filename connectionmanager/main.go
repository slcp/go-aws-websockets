package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/stuartforrest-infinity/websocket-lambda/data"
)

// WebSocketRequest is
type WebSocketRequest struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

// Response is
type Response struct {
	Body string `json:"data"`
}

type snsMessage struct {
	ConnectionID string `json:"connectionId"`
	CallbackURL  string `json:"callbackUrl"`
}

// Handler is
type Handler func(context context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error)

// Handle is
func Handle(context context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("connection handler ran")

	// Store the connectionId in dynamo
	tbname := "websoket-lambda-dev-data"
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	ds := data.NewDataStore(tbname, sess)
	ds.PutConnectionID(data.ConnectionItemData{
		ID: request.RequestContext.ConnectionID,
	}, time.Now())

	m, err := json.Marshal(snsMessage{
		ConnectionID: request.RequestContext.ConnectionID,
		CallbackURL:  fmt.Sprintf("https://%s/%s", request.RequestContext.DomainName, request.RequestContext.Stage),
	})
	if err != nil {
		fmt.Println("marshalling error")
		fmt.Println(err)
	}
	snsclient := sns.New(sess)
	_, err = snsclient.Publish(&sns.PublishInput{
		TopicArn:         aws.String(os.Getenv("TOPIC_ARN")),
		Message:          aws.String(string(m)),
	})
	if err != nil {
		fmt.Println("sns publish error")
		fmt.Println(err)
	}
	// url := fmt.Sprintf("https://%s/%s", request.RequestContext.DomainName, request.RequestContext.Stage)
	// apigw := apigatewaymanagementapi.New(sess, &aws.Config{
	// 	Endpoint: aws.String(url),
	// })
	// conns, err := ds.GetAllConnectionIDs(time.Now())
	// if err != nil {
	// 	fmt.Println("get data error")
	// 	fmt.Println(err)
	// }
	// for _, conn := range conns {
	// 	_, err := apigw.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
	// 		ConnectionId: aws.String(conn.ID),
	// 		Data: []byte(`hello`),
	// 	})
	// 	if err != nil {
	// 		fmt.Println("error")
	// 		fmt.Println(err)
	// 	}
	// }
	// webSocketRequest := &WebSocketRequest{}
	// if err := json.Unmarshal([]byte(request.Body), webSocketRequest); err != nil {
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode:      500,
	// 		IsBase64Encoded: false,
	// 		Body:            "notstringy",
	// 		Headers: map[string]string{
	// 			"Content-Type": "application/json",
	// 		},
	// 	}, nil
	// }
	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	return resp, nil
}

func main() {
	fmt.Println("handler starting")
	lambda.Start(Handle)
}
