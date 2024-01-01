package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	jsonDocument := `
	{
		"name": "kn1ght",
		"age": 20
	}
	`

	var data map[string]interface{} // Declare a map to save JSON docs
	// key is string, value is interface that can hold any value
	json.Unmarshal([]byte(jsonDocument), &data)
	fmt.Printf("name => %v(%v), age => %v(%v)\n",
		data["name"], reflect.TypeOf(data["name"]), data["age"], reflect.TypeOf(data["age"]))

}
