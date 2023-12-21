package main

import (
	"fmt"
	"math"
)

// With using Newton's method,
// get sqrt(x) manually and gradually
func Sqrt(x float64) float64 {
	z := float64(1)
	seq := int(1)
	for ; seq <= 100; seq++ {
		fmt.Printf("Try %d => Value: %g\n", seq, z)
		z -= (math.Pow(z, 2) - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Printf("Got a value of sqrt(2): %g\n", Sqrt(2))
	fmt.Printf("Real value of Sqrt(2):  %g\n", math.Sqrt(2))
}
