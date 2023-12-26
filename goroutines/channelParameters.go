package main

import "fmt"

// This function is only supposed to send data to the channel
// use parameter "chan<-"
func sendToChannel(ch chan<- string) {
	ch <- "DataGoneWild"
	// x := <- ch  // (no!)
}

// This function is only supposed to receive data from the channel
// use parameter "<-chan"
func receiveFromChannel(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
	// ch <- "Data"	// (no!)
}

func main() {
	ch := make(chan string, 1)
	sendToChannel(ch)
	receiveFromChannel(ch)
}
