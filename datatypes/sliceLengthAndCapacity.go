package main

import (
	"fmt"
)

func main() {

	s := []int{2, 3, 5, 7, 11, 13}
	printSlices(s)

	s = s[:0]
	printSlices(s)

	s = s[:4]
	printSlices(s)

	// s = s[2:123]
	s = s[2:]
	printSlices(s)

}

func printSlices(s []int) {
	fmt.Printf("length => %d | capacity => %d | literal => %v\n", len(s), cap(s), s)
}
