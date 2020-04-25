package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// WebSocketRequest is
type WebSocketRequest struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

// Response is
type Response struct {
	Message string `json:"message"`
}

// Handler is
type Handler func(context context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error)

// Handle is
func Handle(context context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("messages handler ran")
	var webSocketRequest WebSocketRequest
	if err := json.Unmarshal([]byte(request.Body), &webSocketRequest); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode:      500,
			IsBase64Encoded: false,
			Body:            "notstringy",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
	fmt.Println(webSocketRequest)
	r := Response{
		Message: "via routeResponseSelectionExpression",
	}
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println("error marshalling response body")
		return events.APIGatewayProxyResponse{
			StatusCode:      500,
			IsBase64Encoded: false,
			Body:            "notstringy",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}
	fmt.Println("returning 200 with body")
	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(b),
	}
	return resp, nil
}

func main() {
	fmt.Println("handler starting")
	lambda.Start(Handle)
}
