package main

import (
	"fmt"
	"time"
)

func main() {

	channelString := make(chan string)
	channelInteger := make(chan int)

	go func(message string) {
		time.Sleep(time.Second * 1)
		channelString <- message
	}("example string message :)")

	go func(value int) {
		time.Sleep(time.Second * 2)
		channelInteger <- value
	}(777)

	for i := 0; i < 2; i++ {
		select {
		case message := <-channelString:
			fmt.Printf("channelString received: %v\n", message)
		case value := <-channelInteger:
			fmt.Printf("channelInteger received: %v\n", value)
		}
	}

}
