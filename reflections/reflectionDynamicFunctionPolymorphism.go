package main

import (
	"fmt"
	"reflect"
)

// It will not provides polymorphism for various datatypes
// But I'd like to use generic functions instead
func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]

	// unsupported case
	if a.Kind() != b.Kind() {
		fmt.Printf("%v and %v has different types.\n", a, b)
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func sumFunctionConnection(functionPointer interface{}) {
	fn := reflect.ValueOf(functionPointer).Elem()
	v := reflect.MakeFunc(fn.Type(), sum)
	fn.Set(v)
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	sumFunctionConnection(&intSum)
	sumFunctionConnection(&floatSum)
	sumFunctionConnection(&stringSum)

	fmt.Println(intSum(7777, 9999))
	fmt.Println(floatSum(2.71828, 1.8284))
	fmt.Println(stringSum("Polymorphism in ", "Golang reflection?"))
}
