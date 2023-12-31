package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// For complex numbers (the input is less than zero(0))
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
