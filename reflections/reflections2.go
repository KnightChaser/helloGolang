package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 1.3

	t := reflect.TypeOf(f)  // save the type information of "f" to "t"
	v := reflect.ValueOf(f) // save the value information of "f" to "v"

	// Extracting type information
	fmt.Printf("name => %v, size => %v, Kind(reflect.Float64) => %v, Kind(reflect.Int64) => %v\n",
		t.Name(), t.Size(), t.Kind() == reflect.Float64, t.Kind() == reflect.Int64)

	// Extracting value information
	fmt.Printf("type => %v, Kind(reflect.Float64) => %v, Kind(reflect.Int64) => %v, value(Float) => %v\n",
		v.Type(), v.Kind() == reflect.Float64, v.Kind() == reflect.Int64, v.Float())

}
