// Reflection is a functionality that grabs the datatype during the runtime
package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num))

	var s string = "hello golang reflect"
	fmt.Println(reflect.TypeOf(s))

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f))

	var data Data = Data{1000, 6666}
	fmt.Println(reflect.TypeOf(data))
}

// The output
// int
// string
// float32
// main.Data
