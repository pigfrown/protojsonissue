package main

import (
	"fmt"

	"github.com/pigfrown/protojsonissue"
)

func main() {
	convertTwoMessage()
}

func convertOneMessage() {
	msg := &protojsonissue.ExampleMessage{
		AString:   "somestring",
		SomeBytes: []byte("somebytes"),
		ANumber:   23,
	}

	fmt.Println(protojsonissue.ExampleMessageToJSON(msg))
}

func convertTwoMessage() {
	msg := &protojsonissue.ExampleMessage{
		AString:   "somestring",
		SomeBytes: []byte("somebytes"),
		ANumber:   23,
	}

	fmt.Println(protojsonissue.ExampleMessageToJSON(msg))
	fmt.Println(protojsonissue.ExampleMessageToJSON(msg))
}
