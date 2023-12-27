package main

import "fmt"

// With only using channel, implement integer addition

// This function returns the "only integer receiving channel"
func num(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		out <- b
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
		for i := range c {
			r += i
		}
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
