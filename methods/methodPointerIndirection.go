package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// Pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	// Golang automatically interprets v.Scale(2) as (&v).Scale(5)
	// even though v is not a pointer, it's an immediate value because
	// the method has a pointer receiver
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
