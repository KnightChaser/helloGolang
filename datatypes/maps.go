package main

import (
	"fmt"
)

type Vertex struct {
	Latitude, Longitude float64
}

var m map[string]Vertex

func main() {
	fmt.Println(m) // just a shell, nothing. (map[])
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.234, -12.1239,
	}

	fmt.Println(m)              // map[Bell Labs:{40.234 -12.1239}]
	fmt.Println(m["Bell Labs"]) // {40.234 -12.1239}
}
