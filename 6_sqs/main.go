package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

type Person struct {
	Firstname string
	Lastname  string
}

var persons = []Person{
	Person{"Fyodor", "Dostoevsky"},
	Person{"William", "Shakespeare"},
	Person{"Charles", "Dickens"},
	Person{"Leo", "Tolstoy"},
	Person{"Mark", "Twain"},
	Person{"George", "Orwell"},
	Person{"Ernest", "Hemingway"},
}

// from environment
var queueURL string
var groupID string

var processed []string

type Request struct {
	Action string `json:"action"` // populate, process or get
}

func handleRequest(request Request) (string, error) {
	switch request.Action {
	case "populate":
		return Populate(&queueURL, &groupID, persons)
	case "process":
		return Process(&queueURL, &processed)
	default:
		return strings.Join(processed, ","), nil
	}
}

func main() {
	queueURL = os.Getenv("QUEUE_URL")
	groupID = os.Getenv("QUEUE_GROUP_ID")
	processed = []string{}
	lambda.Start(handleRequest)
}
