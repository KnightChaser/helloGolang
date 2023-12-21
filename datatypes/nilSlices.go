package main

import (
	"fmt"
)

func main() {
	// nil stands for zero or "uninitialized" in the data object in golang.
	// It's length and capacity will be zero as a default
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Printf("The data %v is currently nil!\n", s)
	}
}
