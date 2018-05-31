package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"strings"
)

func Process(queueURL *string, processed *[]string) (string, error) {

	svc := sqs.New(session.Must(session.NewSession()))

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(2),
	})
	if err != nil {
		return "", err
	}

	nbMessages := len(result.Messages)
	if nbMessages == 0 {
		return "OK", nil
	}

	for i := 0; i < nbMessages; i += 1 {

		var p Person
		if err := json.Unmarshal([]byte(*result.Messages[i].Body), &p); err != nil {
			panic(err)
		}

		*processed = append(*processed, strings.Join([]string{p.Firstname, p.Lastname}, " "))

		if _, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      queueURL,
			ReceiptHandle: result.Messages[i].ReceiptHandle,
		}); err != nil {
			return "", err
		}

	}

	return "OK", nil
}
