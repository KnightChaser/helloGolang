// A conditional variable in goroutine is used to wake a sleeping object or multiple sleeping objects up
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)
	// sync.NewCond(mutex).Wait(): pause goroutine process and wait
	// sync.NewCond(mutex).Signal(): wake up only one sleeping goroutine
	// sync.NewCond(mutex).Broadcast(): wake up every sleeping goroutine
	var goroutineQty int = 5
	c := make(chan bool, goroutineQty) // Create an async boolean channel

	for i := 0; i < goroutineQty; i++ {
		go func(n int) {
			mutex.Lock() // Lock mutex
			c <- true
			fmt.Printf("Wait begins... goroutine#%d\n", n)
			cond.Wait() // wait conditional variable
			fmt.Printf("Wait ends..... goroutine#%d\n", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < goroutineQty; i++ {
		<-c // Pops every boolean value from the boolean channel c
	}

	// Protect the cond.Broadcast() process
	// Because the process will wake up at once, the procedure of waking up is not guaranteed unlike sync.NeweCond(mutex).Signal()
	mutex.Lock()
	fmt.Println("Broadcast (wake up, every goroutine!)")
	cond.Broadcast()
	mutex.Unlock()

	fmt.Scanln()
}
