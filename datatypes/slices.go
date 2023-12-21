package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice a[low:high] like python list slicing.
	// However, slice is like a pointer to the certain range of the target array
	var s []int = primes[1:4]
	fmt.Println(s)
}
