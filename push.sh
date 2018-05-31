#!/bin/bash

project="$1"
[ $project = "" ] && echo "Project name is required." && exit 1

export GOOS=linux
export GOARCH=amd64

cd $project

go build -ldflags '-w -s' -o main \
  && zip -X main.zip * \
  && aws lambda update-function-code --function-name=test \
      --zip-file=fileb:///home/ec2-user/environment/$project/main.zip \
  && rm -f main main.zip
