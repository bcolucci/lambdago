package main

import (
	"encoding/base64"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lbd "github.com/aws/aws-sdk-go/service/lambda"
	"time"
)

func handleRequest() (interface{}, error) {

	ctx := Stringify(map[string]interface{}{
		"from": "3_invoke",
		"time": time.Now(),
	})

	event := Stringify(map[string]interface{}{
		"sid":      "Charlie",
		"Lastname": "Colucci",
		"Age":      3,
	})

	svc := lbd.New(session.Must(session.NewSession()))

	result, err := svc.Invoke(&lbd.InvokeInput{
		ClientContext: aws.String(base64.StdEncoding.EncodeToString(ctx)),
		FunctionName:  aws.String("ping"),
		Payload:       event,
	})
	if err != nil {
		return nil, err
	}

	return Parse(result.Payload), nil
}

func main() {
	lambda.Start(handleRequest)
}
