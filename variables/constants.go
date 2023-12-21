package main

import (
	"fmt"
)

// Const values(constant values) cannot be declared with using ":=" keyword
const Pi = 3.14 // const

func getCircleSpace(x float64) float64 {
	return float64(x) * float64(x) * Pi
}

func main() {
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	radius := 123.45678
	fmt.Printf("The space of circle with raidus %v has a space of %v\n", radius, getCircleSpace(radius))

	const Truth = true
	fmt.Println("Go grammer weird? =>", Truth)
}
