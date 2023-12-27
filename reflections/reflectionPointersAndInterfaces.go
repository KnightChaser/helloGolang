package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Create an integer pointer and apply 1 to it
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	// fmt.Println(reflect.ValueOf(a).Int())		// panic: reflect: call of reflect.Value.Int on ptr Value
	fmt.Println(reflect.ValueOf(a).Elem())
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // because it's pointer, need to access the real value by Elem()

	var b interface{}
	b = 1
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(b).Int())
	// fmt.Println(reflect.ValueOf(b).Elem())		// panic: reflect: call of reflect.Value.Elem on int Value

}
