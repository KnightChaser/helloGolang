package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	// waiting multiple channels to diversify functionalities
	for {
		select {
		// In case of data receiving through a channel
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			// If there's default in select,
			// Go runtime just executes default statement even if
			// the channels aren't prepared yet.
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 25; i++ {
			// Inserting data to the channel
			fmt.Println(<-c)
		}
		// Afer some certain tasks, send 0 to quite channel to terminate fibonacci()
		quit <- 0
	}()
	// The channel is prepared. Run fibonacci() function.
	fibonacci(c, quit)
}
