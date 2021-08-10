package main

import (
	"fmt"

	"github.com/pigfrown/protojsonissue"
)

func main() {

	msg := &protojsonissue.ExampleMessage{
		AString:   "somestring",
		SomeBytes: []byte("somebytes"),
		ANumber:   23,
	}

	fmt.Println(protojsonissue.ExampleMessageToJSON(msg))
	fmt.Println(protojsonissue.ExampleMessageToJSON(msg))
}
