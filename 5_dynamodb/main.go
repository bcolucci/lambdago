package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	SID string `json:"sid"`
}

func handleRequest(request Request) (DemoSession, error) {
	var session *DemoSession
	if request.SID != "" {
		session = LoadDemoSession(request.SID)
	}
	if session == nil {
		session = NewDemoSession()
	}
	session.Views += 1
	session.Save()
	return *session, nil
}

func main() {
	lambda.Start(handleRequest)
}
