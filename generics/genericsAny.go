package main

import (
	"fmt"
)

// List represents a singly-linked list that holds values of any type by using [T any]
func print[T any](input T) {
	fmt.Printf("Value: %v, Type: %T\n", input, input)
}

func main() {
	var (
		a int     = 12345678
		b float64 = 2.71828
		c string  = "Hello golang generics"
	)
	print(a)
	print(b)
	print(c)
}
