package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	say("Synchronous")

	go say("Asynchronous #1")
	go say("Asynchronous #2")
	go say("Asynchronous #3")

	time.Sleep(time.Second * 3)
}
