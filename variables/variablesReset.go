package main

import (
	"fmt"
)

var i, j int = 1, 2 // Declare and reset multiple variables at once

func main() {
	var c, python, java = true, false, "No!" // type guessing according to the immediates
	fmt.Println(c, python, java)
}
