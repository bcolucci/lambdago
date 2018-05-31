#!/bin/bash

API_URL="https://j7sxohq3s7.execute-api.eu-west-1.amazonaws.com/test"

curl -X GET $API_URL/part1/?q1=param1; echo
curl -X POST -d '{"username":"brice","password":"my_secret"}' $API_URL/; echo
