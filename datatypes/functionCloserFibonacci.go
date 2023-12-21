package main

import (
	"fmt"
)

func fibonacci() func() int {
	// First term of the sequence
	a, b := 0, 1
	return func() int {
		result := a
		// Because it's a closer, the local variables won't be vanished away
		a, b = b, a+b // Inductive(recursive) def of the sequence
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
