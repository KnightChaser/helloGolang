package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"knightchaser",
		"krisse",
		"klojure",
		"krlierr",
	}
	fmt.Println(names)

	// Because slice is like a pointer that indicates
	// some certain range of an array. If you modify the slice,
	// the target part of slice will be changed too
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "???"
	fmt.Println(a, b)
	fmt.Println(names)
}
