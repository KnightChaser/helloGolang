package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	// The method is completely same with "./goroutineAtomicYes2.go", but since Go 1.19
	// utilizing sync/atomic pkg to every possible area
	var new atomic.Int64
	new.Store(123)
	var old = new.Swap(789)
	fmt.Printf("new => %d, old => %d\n", new.Load(), old)

	swapped := new.CompareAndSwap(123, 456)
	fmt.Printf("swapped => %v\n", swapped)
	fmt.Printf("new     => %d\n", new.Load())

	swapped = new.CompareAndSwap(789, 456)
	fmt.Printf("swapped => %v\n", swapped)
	fmt.Printf("new     => %d\n", new.Load())

}
