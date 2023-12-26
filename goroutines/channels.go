package main

import "fmt"

func main() {
	ch := make(chan int) // Create an integer channel named "ch" for integer data transfer

	go func() {
		ch <- 123 // send 123 via the channel
	}()

	i := <-ch // Receive a data(integer 123) from a channel named "ch"

	fmt.Println(i)
}
