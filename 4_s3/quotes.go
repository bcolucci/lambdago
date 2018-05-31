package main

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"math/rand"
	"strings"
	"time"
)

var quotes []string

func lazyLoad() {

	svc := s3.New(session.Must(session.NewSession()))

	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("brice-lambdago"),
		Key:    aws.String("quotes.txt"),
	})
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, result.Body); err != nil {
		panic(err)
	}

	quotes = strings.Split(string(buf.Bytes()), "\n")
}

func RandomQuote() string {
	if len(quotes) == 0 {
		lazyLoad()
	}
	rand.Seed(time.Now().Unix())
	return quotes[rand.Int()%len(quotes)]
}
