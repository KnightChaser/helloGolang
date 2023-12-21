package main

import (
	"fmt"
)

func main() {
	var a [2]string // => char* a[2]
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Print element by element
	fmt.Println(a)          // Print the array itself at once

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // Print the array itself at once
}
