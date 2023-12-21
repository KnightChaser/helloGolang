package main

import (
	"fmt"
)

type Vertex struct {
	Latitude, Longitude float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])
	fmt.Println(m["invalid??"]) // no data, just returning {0 0}
}
