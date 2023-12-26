package main

import "fmt"

func main() {
	// Channels can be buffered with length. They work like queues
	ch := make(chan int, 2)
	ch <- 1 // Nothing receives the value... but the channel is "buffered", so it can store data up to the limit
	ch <- 2
	fmt.Println(<-ch) // receives 1
	fmt.Println(<-ch) // receives 2
	// fmt.Println(<-ch)  	// runtime error "all goroutines are asleep"
	// because there is nothing in the channel, so can't print data from the channel anymore
}
