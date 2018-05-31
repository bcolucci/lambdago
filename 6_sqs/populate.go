package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Populate(queueURL, groupID *string, persons []Person) (string, error) {
	svc := sqs.New(session.Must(session.NewSession()))
	for i := 0; i < len(persons); i += 1 {
		json, _ := json.Marshal(persons[i])
		body := string(json)
		if _, err := svc.SendMessage(&sqs.SendMessageInput{
			QueueUrl:       queueURL,
			MessageGroupId: groupID,
			MessageBody:    &body,
		}); err != nil {
			return "", err
		}
	}
	return "OK", nil
}
