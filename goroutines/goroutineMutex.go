package main

import (
	"fmt"
	"sync"
)

// A safe counter is safe to use "concurrently"
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Increase function increases the counter for the given key
func (sc *SafeCounter) Increase(key string, done chan bool) {
	sc.mu.Lock()
	// critical section!
	sc.v[key]++
	sc.mu.Unlock()

	// signal that the current operation got accomplished
	done <- true
}

// Value returns the current value of the counter for the given key
func (sc *SafeCounter) GetValueFromSafeCounter(key string) int {
	sc.mu.Lock()
	// critical section!
	defer sc.mu.Unlock() // so the unlock will be done at that moment of the function termination
	return sc.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	done := make(chan bool)
	qtyGoroutines := 7777

	for i := 0; i < qtyGoroutines; i++ {
		go c.Increase("some_key", done)
		go c.Increase("tacotaco", done)
	}

	// Wait for all goroutines to finish
	for i := 0; i < qtyGoroutines*2; i++ {
		<-done
	}

	fmt.Println(c.GetValueFromSafeCounter("some_key"))
	fmt.Println(c.GetValueFromSafeCounter(("tacotaco")))
}
