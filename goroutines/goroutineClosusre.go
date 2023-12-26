package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var waitGroup sync.WaitGroup
	s := "Hello, world!"
	goroutineExecutionQty := 100

	for i := 0; i < goroutineExecutionQty; i++ {
		waitGroup.Add(1)
		go func(n int) {
			defer waitGroup.Done()
			fmt.Println(s, n)
		}(i) // if you don't pass the iteration variable as a parameter,
		// the goroutines will be executed after iteration, so i = 100 will be executed 100 times.
	}

	waitGroup.Wait()

}
