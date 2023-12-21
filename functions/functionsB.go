// A function can return multiple values
package main

import (
	"fmt"
)

// swap function, which reverses the order of the given parameters to the function
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// fmt.Println(swap("Hello", "World"))
	a, b := swap("Hello", "World") // Allocate the values from the function
	fmt.Printf("%s %s\n", a, b)
}
