package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var data atomic.Int32
	var wg sync.WaitGroup

	// Add 1 to the variable named "data" 2,000 times atomically
	for i := 0; i < 1000; i++ {
		wg.Add(1) // Add the value to the target variable using sync/atomic official method
		go func() {
			data.Add(1)
			wg.Done()
		}()
	}

	// Subtract 1(actually, adding -1) to the variable named "data" 2,000 times atomically
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			data.Add(-1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data.Load()) // Retrieve the value using sync/atomic official method

}
