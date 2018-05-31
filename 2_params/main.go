package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func handleRequest(request Request) (Response, error) {
	var name = "world"
	if request.Name != "" {
		name = request.Name
	}
	return Response{fmt.Sprintf("Hello, %s!", name)}, nil
}

func main() {
	lambda.Start(handleRequest)
}
