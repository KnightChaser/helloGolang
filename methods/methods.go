package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// The method(~function) Abs() has a Vertex type receiver named v
// We can call 'em as a receiver functions
// Receiver functions gets parameters even prior than a function name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // such like class method (However, Golang doesn't have classes officially)
}
