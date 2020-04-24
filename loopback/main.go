package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

// Response is
type Response struct {
	Body string `json:"data"`
}

type snsMessage struct {
	ConnectionID string `json:"connectionId"`
	CallbackURL  string `json:"callbackUrl"`
}

// Handler is
type Handler func(context context.Context, records events.SNSEvent) error

// Handle is
func Handle(context context.Context, records events.SNSEvent) error {
	fmt.Println("loopback handler ran")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	for _, r := range records.Records {
		var m snsMessage
		err := json.Unmarshal([]byte(r.SNS.Message), &m)
		if err != nil {
			fmt.Println("sns unmarshal error")
			fmt.Println(err)
		}
		url := fmt.Sprintf(m.CallbackURL)
		apigw := apigatewaymanagementapi.New(sess, &aws.Config{
			Endpoint: aws.String(url),
		})
		_, err = apigw.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(m.ConnectionID),
			Data:         []byte(`hello`),
		})
		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
		}
	}
	return nil
}

func main() {
	fmt.Println("handler starting")
	lambda.Start(Handle)
}
