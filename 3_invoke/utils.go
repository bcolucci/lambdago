package main

import (
	"encoding/json"
)

func Stringify(obj interface{}) []byte {
	json, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return json
}

func Parse(bytes []byte) interface{} {
	var obj interface{}
	if err := json.Unmarshal(bytes, &obj); err != nil {
		panic(err)
	}
	return obj
}
