package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handleRequest(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	json, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{Body: string(json), StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(handleRequest)
}
