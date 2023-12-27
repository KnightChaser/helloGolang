package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	// close the channel
	close(ch)

	// Print every existing data in the channel until it exhausts
	for i := range ch {
		fmt.Println(i)
	}
}
