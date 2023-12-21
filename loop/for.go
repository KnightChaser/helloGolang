package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("The sum from %v to %v is %v\n", 1, 10, sum)
}
