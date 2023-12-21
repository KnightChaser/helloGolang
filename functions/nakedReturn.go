package main

import (
	"fmt"
)

// Naked return
//
//			|---- parameter
//	 					|---- designate returning values' name and types
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // implicitly return the value, (x, y int) will be returned as declared
}

func main() {
	fmt.Println(split(17))
}
