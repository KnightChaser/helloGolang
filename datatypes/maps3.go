package main

import "fmt"

type Position struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Position)
	m["PositionA"] = Position{12.3456, 23.4995}
	m["PositionB"] = Position{58.5555, 47.3334}
	fmt.Println(m)
	fmt.Println(m["PositionA"])
}
