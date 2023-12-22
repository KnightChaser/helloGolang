package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// If (v *Vertex) => 50 returns
// If (v Vertex)  =>  5 returns
// It's a pointer receiver. It modififes the target object pointed by a pointer, so it's not necessarily required to specify returning values like others.
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
