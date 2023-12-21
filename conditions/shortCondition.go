package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g > %g\n", v, lim)
		return lim
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 10),    // Prints 9 because 3^2 < 10
		pow(3, 3, 20),    // Prints 20 because 3^3 > 20
		pow(10, 3, 1024), // Prints 1000 because 10^3 < 1024
	)
}
