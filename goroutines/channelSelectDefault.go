package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-tick:
				fmt.Println("tick!")
			case <-boom:
				fmt.Println("boom!")
			case <-stop:
				fmt.Println("Terminated!!")
				break
			default:
				// The default case in a select is run if no other case is ready.
				fmt.Println(".")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Millisecond * 2000)
	stop <- true
}
