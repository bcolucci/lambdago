package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strconv"
)

type DemoSession struct {
	SID   string `json:"sid"`
	Views int    `json:"nb_views"`
}

func NewDemoSession() *DemoSession {
	sid := NewSID(32)
	return &DemoSession{SID: sid, Views: 0}
}

func LoadDemoSession(sid string) *DemoSession {

	svc := dynamodb.New(session.Must(session.NewSession()))

	var err error
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("session"),
		Key: map[string]*dynamodb.AttributeValue{
			"sid": {
				S: aws.String(sid),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	item := DemoSession{}
	if err = dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		panic(err)
	}

	return &item
}

func (sess *DemoSession) Save() {

	svc := dynamodb.New(session.Must(session.NewSession()))

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("session"),
		Key: map[string]*dynamodb.AttributeValue{
			"sid": {
				S: aws.String(sess.SID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v": {
				N: aws.String(strconv.Itoa(sess.Views)),
			},
		},
		UpdateExpression: aws.String("set nb_views = :v"),
	}

	if _, err := svc.UpdateItem(input); err != nil {
		panic(err)
	}

}
