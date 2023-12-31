package main

import (
	"fmt"
)

const (
	// Undetermined type of constants automatically takes its type in the given context
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("type %T => value %v\n", needInt(Small), needInt(Small))
	fmt.Printf("type %T => value %v\n", needFloat(Small), needFloat(Small))
	fmt.Printf("type %T => value %v\n", needFloat(Big), needFloat(Big))
	// fmt.Println(needInt(Big))		// overflow
}
