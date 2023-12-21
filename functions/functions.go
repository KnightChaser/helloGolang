package main

import (
	"fmt"
	"time"
)

func add(x int, y int) int {
	return x + y
}

// Because all parameters of the given function are the same as integer,
// We can abbreviate the form from "x int, y int" to "x, y int"
func multiply(x, y int) int {
	return x * y
}

func printCurrentTime() {
	fmt.Printf("The current time is %s\n", time.Now())
}

func main() {
	fmt.Println(add(100, 200))
	fmt.Println(multiply(100, 200))
	printCurrentTime()
}
