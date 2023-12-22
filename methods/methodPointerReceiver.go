// Generally pointer receivers are better than value receivers.
// It's the same why C/C++ support pointers
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Why method pointer recivers are good options in general?
// - Avoid copy problem (creating new objs)
// - Able to modify the values that the reciver directs (Because it's a pointer)
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
