package main

import (
	"fmt"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		result := y
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Printf("f(%v) => %v\n", i, f())
	}
}
