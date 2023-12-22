package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// An empty interface => interface{}
// It may hold values of any type. (Every type implements at least zero methods, meaning {}.)
// Every interfaces are used by code that handles values of unknown type,. For example,
// fmt.Print takes any number of arguments of type interface{}.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	var i interface{}

	// not initialized yet. Thus, <nil> will be.
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = Vertex{12.345, -76.543}
	describe(i)

}
