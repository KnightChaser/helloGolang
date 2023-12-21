package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("time.Now() => %v\n", t)

	// switch phrase with no condition, it workss switch-true.
	// It might be a very clear way to write very long if-else conditional chain
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
