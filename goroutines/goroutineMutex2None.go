package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// during the parallel tasks, utilize every possible CPU
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := []int{}
	goroutineExecutionQty := 1000
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		// by using defer, set WaitGroup as done right before the function finishes working
		defer wg.Done()
		for i := 0; i < goroutineExecutionQty; i++ {
			data = append(data, i)
			// Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine,
			// so execution resumes automatically.
			runtime.Gosched()
		}
	}()

	// another goroutine trying to append values
	go func() {
		defer wg.Done()
		for i := 0; i < goroutineExecutionQty; i++ {
			data = append(data, i)
			runtime.Gosched()
		}
	}()

	// Wait until every working group(goroutines) finishes working
	wg.Wait()
	fmt.Println(len(data))
}
