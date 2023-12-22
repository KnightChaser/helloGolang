// A type switch is a construct that permits several type assertions in series
// So you can proceed the processes case-by-case, according to its value type, with using switch-case sentences
package main

import (
	"fmt"
	"math"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("(integer) Twice %v is %v\n", v, math.Pow(float64(v), 2))
	case string:
		fmt.Printf("(string) %q is %v bytes long...\n", v, len(v))
	case bool:
		fmt.Printf("(boolean) %v is %v!\n", v, v)
	default:
		fmt.Printf("(%T) Something else that I don't know yet :( => %v\n", v, v)
	}
}

type Person struct {
	Name string
	Age  uint
}

func main() {
	do(21)
	do("Hello World")
	do(true)
	do(Person{"Klojure", 20})
}
