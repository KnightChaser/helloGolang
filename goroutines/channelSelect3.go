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
			c1 <- 777
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello Golang goroutines!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Just created a 333 millisecond long time tick
	go func() {
		for {
			<-time.After(333 * time.Millisecond)
			fmt.Printf("A 333 millisecond time tick just ticked!\n")
		}
	}()

	go func() {
		for {
			select {
			case integerValue := <-c1:
				fmt.Printf("Channel 1 says: %d\n", integerValue)
			case stringValue := <-c2:
				fmt.Printf("Channel 2 says: %s\n", stringValue)
			}
		}
	}()

	// Let's see what happens for 10 seconds
	time.Sleep(10 * time.Second)
}

// The output will be...
// Channel 1 says: 777
// Channel 2 says: Hello Golang goroutines!
// Channel 1 says: 777
// Channel 1 says: 777
// Channel 1 says: 777
// Channel 1 says: 777
// Channel 2 says: Hello Golang goroutines!
// Channel 1 says: 777
// Channel 1 says: 777
// Channel 1 says: 777
// ...
