package main

import (
	"fmt"
)

func main() {

	s := []int{2, 3, 5, 7, 11, 13}

	// default value at the slice(slice default) will be the starting or ending point
	fmt.Println(s[1:4])
	fmt.Println(s[:2])
	fmt.Println(s[1:])

}
