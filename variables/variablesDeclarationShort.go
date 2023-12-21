package main

import (
	"fmt"
)

func main() {
	var i, j int = 1, 2
	k := 3 // shorter shorter
	// At the outside of the function, everything must be starting with keywords,
	// so declaring variables will only be "var i int = 1", something like this.
	c, python, java := true, false, "No!"
	fmt.Println(i, j, k, c, python, java)
}
