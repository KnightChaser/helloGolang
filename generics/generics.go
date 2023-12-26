// Generic is a function that accepts multiple type of variables, general function
// such like polymorphism in the modern C++
package main

import "fmt"

// Index rreturns the index of x in s, or -1 if not found
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constratin, so we can use == here
		if v == x {
			// return its index
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "bazz", "lmfao"}
	fmt.Println(Index(ss, "hello world"))
}
