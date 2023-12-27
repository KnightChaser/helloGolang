package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	// Just adding, subtracing operation is not a safe operation because it isn't atomic

	// Add 1 to the variable "data" 2,000 times
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data += 1
		}()
	}

	// Subtract 1 to the varaible "data" 2,000 times
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data -= 1
		}()
	}

	wg.Wait()
	fmt.Println(data) // The vale was supposed to be 0, but it isn't actually.

}
