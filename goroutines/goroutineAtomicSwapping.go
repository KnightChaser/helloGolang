package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var new int64 = 123
	var old = atomic.SwapInt64(&new, 789)
	fmt.Printf("new => %d, old => %d\n", new, old)

	// The atomic.CompareAndSwapInt64() function checks
	// if the current value at the memory location pointed to by &new is equal to the expected value (123).
	// If it is, it atomically updates the value to the new value (456) and returns true.
	// Otherwise, it returns false.
	swapped := atomic.CompareAndSwapInt64(&new, 123, 456)
	fmt.Printf("swapped => %v\n", swapped)
	fmt.Printf("new     => %d\n", new)

	// *(&new) is currently 789. The second parameter of atomic.CompareAndSwapInt64() is 789,
	// thus it returns true and *(&new)'s value will be swapped into 456
	swapped = atomic.CompareAndSwapInt64(&new, 789, 456)
	fmt.Printf("swapped => %v\n", swapped)
	fmt.Printf("new     => %d\n", new)

}
