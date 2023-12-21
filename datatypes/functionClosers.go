package main

import (
	"fmt"
)

// The return value is the anonymous functio! (function in function, AKA nested function smth)
func adder() func(int) int {
	sum := 0
	// closer in function, create an anonymous function and declare it. In this case it's for returning
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
