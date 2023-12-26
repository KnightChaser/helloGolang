// Channels are a typed conduit through which you can send and receive avalues with the channel operator, "<-".

package main

import "fmt"

//					|--- meaning "channel" actually
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel(chan) c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 8}
	c := make(chan int) // create a channel for integer types
	// must use make() and be created before usage

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive a data from channel c from func sum

	fmt.Println(x, y, x+y)
}
