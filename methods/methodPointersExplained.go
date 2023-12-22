package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func Scale(v Vertex, f float64) Vertex {
	v.X *= f
	v.Y *= f
	return v
}

func main() {
	v := Vertex{3, 4}
	v = Scale(v, 10)
	fmt.Println(Abs(v))
}
