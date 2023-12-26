// A goroutine is a lightweight thread that is managed by Go runtime
// It runs Golang instructions asynchronously, so it used to run multiple codes parallely.
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// Starts a new goroutine with using keyword "go"
	go say("world")
	say("hello")
}
