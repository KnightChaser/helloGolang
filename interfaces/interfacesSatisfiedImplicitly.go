package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T "implemens" the interface I,
// but we don't need to explicitly declare that it does so.
// --> "type I interface" defined a method(~function) T
func (t T) M() {
	fmt.Println(t.S) // Print the string which is a property of the object T (type T struct)
}

func main() {
	var i I = T{"Hello Golang interfaces"}
	i.M() // By calling, it calls fmt.Println(t.S) by func(t T) M()
}
