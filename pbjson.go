package protojsonissue

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
)

func ExampleMessageToJSON(msg *ExampleMessage) string {
	json, err := protojson.Marshal(msg)

	if err != nil {
		fmt.Printf("Could not marshal message to JSON : %v", err)
		return ""
	}
	return string(json)
}
