package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

// This method means type T "implemens" the interface I,
// but we don't need to explicitly declare that it does so.
// --> "type I interface" defined a method(~function) T
func (t T) M() {
	fmt.Println(t.S) // Print the string which is a property of the object T (type T struct)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I = T{"Hello Golang interfaces"}
	describe(i)
	i.M() // By calling, it calls fmt.Println(t.S) by func(t T) M()

	i = F(math.Pi)
	describe(i)
	i.M()
}
