// Observe fmt.Printf() execution order carefully!!!!

package main

import (
	"fmt"
	"time"
)

// With only using channel, implement integer addition

// This function returns the "only integer receiving channel"
func num(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		fmt.Printf("Received(a): %d\n", a)
		time.Sleep(10000 * time.Second)
		out <- b
		fmt.Printf("Received(b): %d\n", b)
		close(out) // close the channel so no one will be able to write on this channel
	}()
	return out
}

// This function receives the "only integer receiving channel"
// This function returns the "only integer sending channel"
func sum(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		temp1 := <-c
		fmt.Printf("Received(temp1): %d\n", temp1)
		temp2 := <-c
		fmt.Printf("Received(temp2): %d\n", temp2)
		r += temp1
		r += temp2
		out <- r
	}()
	return out
}

func main() {
	x := 1234
	y := 3456
	c := num(x, y)
	out := sum(c)
	fmt.Printf("%d + %d = %d\n", x, y, <-out)
}
