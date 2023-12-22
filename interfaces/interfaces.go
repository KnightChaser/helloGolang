package main

import (
	"fmt"
	"math"
)

// Interfaces are like a struct of methods.
// Interfaces define method prototypes that the type(type ... interface) should define
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2) // default math const of sqrt(2)
	v := Vertex{3, 4}
	a = f  // a MyFloat implements Abser (acceptable, because Go automatically changes it suitable for pointer)
	a = &v // a *Vertex implements Abser (okay and expected, because &v is for pointer and Abser(type of a)'s method Abs() is accepting pointer)

	// Error happens
	// a = v // v is a Vertex(not *Vertex) and doens't implement Abser (the only receiver is just *Vertex)
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}
