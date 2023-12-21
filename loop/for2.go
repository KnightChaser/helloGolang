package main

import (
	"fmt"
)

func main() {
	sum := 1
	// You can omit initialization and afterward sentences
	// Thus, for in Golang can be while() in C/C++ and many others
	for sum <= 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
