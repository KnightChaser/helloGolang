package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := make(map[string]interface{}) // key(string) - value(any(interface{}))
	data["name"] = "kn1ght"
	data["age"] = 20
	data["language"] = "Golang"

	document, _ := json.Marshal(data)
	fmt.Println(string(document)) // the original data is byte stream, convert it to string

}
