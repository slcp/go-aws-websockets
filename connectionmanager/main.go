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
	Body string `json:"data"`
}

// Handler is
type Handler func(context context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error)

// Handle is
func Handle(context context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("connection handler ran")
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
	// context.Done()
	// body, err := json.Marshal(Response{
	// 	Body: "stringy",
	// })
	// if err != nil {
	// 	fmt.Println("errorr mashal")
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode: 400,
	// 		Body:       string(body),
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
	// log := logger.Entry("main", "main")

	// var cfg configuration
	// err := config.NewLoader(config.NewEnvSource()).Load(&cfg)
	// if err != nil {
	// 	log.WithError(err).Error("error loading configuration")
	// 	os.Exit(-1)
	// }
	// if err != nil {
	// 	log.WithError(err).Error("error getting google root certs")
	// 	os.Exit(-1)
	// }
	fmt.Println("handler starting")
	lambda.Start(Handle)
}
