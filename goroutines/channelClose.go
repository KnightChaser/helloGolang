package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	close(ch)

	// You can't send to that channel because it's now closed.
	// Channel closing is not mandatory, but you may do this if you have to clarify when the channel is closed
	// Since now, only receiving remaining data in the channel is only available.
	// (e.g. loop statement, to let it know when to stop)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	if _, success := <-ch; !success {
		fmt.Println("No data...")
	}

}
