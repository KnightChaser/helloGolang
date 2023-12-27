package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			// A goroutine fetches an integer value from the channel c1
			i := <-c1
			fmt.Printf("Channel 1 says: %v\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello goroutines"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case c1 <- 10: // every time send an integer valule to channel c1
			case s := <-c2: // if there is a value arrived to the channel c2, apply it to the var s
				fmt.Printf("Channel 2 says: %v\n", s)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
