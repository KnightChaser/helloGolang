package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{10, 20}
	v.X = 15
	fmt.Printf("Vertex struct v => %v\n", v)
	fmt.Printf("v.X => %v | v.Y => %v\n", v.X, v.Y)
}
