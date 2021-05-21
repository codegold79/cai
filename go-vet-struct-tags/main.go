/*
	Structs can be used to (un)marshal json data. What happens with `go vet`
	when a struct has both json and non-json information?
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type mixed struct {
		JsonField    string `json:"jfield"`
		nonJsonField string
	}

	jsonBytes := []byte(`{"jfield":"some json info"}`)

	var mixedData mixed
	json.Unmarshal(jsonBytes, &mixedData)
	fmt.Printf("%#v\n", mixedData)

	mixedData.nonJsonField = "data not from json"
	fmt.Printf("%#v\n", mixedData)
}
