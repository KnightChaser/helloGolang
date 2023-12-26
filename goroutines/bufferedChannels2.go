package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	buffer := make(chan bool, 2)
	bufferFillCount := 10

	go func() {
		for i := 0; i < bufferFillCount; i++ {
			buffer <- true // send true to the buffered(3) channel
			fmt.Printf("Goroutine buffer fill: %d\n", i)
		}
	}()

	for i := 0; i < bufferFillCount; i++ {
		<-buffer // Eject values in the buffer
		fmt.Printf("Goroutine buffer pop: %d\n", i)
	}
}

// wtf it doesn't work
