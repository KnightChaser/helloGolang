package main

import (
	"fmt"
)

func fibonacci(n int, c chan int64) {
	var x int64 = 0
	var y int64 = 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int64, 50)
	go fibonacci(cap(c), c)
	// Exhaust every element in the channel c
	for i := range c {
		fmt.Println(i)
	}
}
