// Reflection can be used to dynamic creation of Golang functions

package main

import (
	"fmt"
	"reflect"
)

// The parameter(s) and return value(s) of the reflection function
// should be []reflect.Value
func h(args []reflect.Value) []reflect.Value {
	fmt.Println("Hello Golang reflection function")
	return nil
}

func main() {
	var hello func()                     // Create a variable to store a function
	fn := reflect.ValueOf(&hello).Elem() // Pass var hello's addr and retrieve value info by Elem()
	v := reflect.MakeFunc(fn.Type(), h)  // Create h's function information
	fn.Set(v)
	hello()
}
