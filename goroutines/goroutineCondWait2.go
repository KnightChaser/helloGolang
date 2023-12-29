package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	cond    *sync.Cond
	counter int
	wg      sync.WaitGroup
)

func main() {

	// Initialize the condition variable with a mutex
	cond = sync.NewCond(&mutex)

	// Launch two goroutines
	wg.Add(2)
	go waitForCondition("Goroutine 1")
	go waitForCondition("Goroutine 2")

	// Simulate a condition change after some time
	time.Sleep(1 * time.Second)
	signalCondition()

	time.Sleep(1 * time.Second)
	signalCondition()

	// Wait for goroutines to finish
	wg.Wait()

}

func waitForCondition(name string) {

	defer wg.Done()

	// Conditional variables should be protected in Mutex
	mutex.Lock()

	fmt.Printf("%s is waiting...\n", name)
	cond.Wait()
	fmt.Printf("%s received the signal. Counter: %d\n", name, counter)

	mutex.Unlock()
}

func signalCondition() {

	// Conditional variables should be protected in Mutex
	mutex.Lock()

	counter++
	fmt.Printf("Condition met! Counter: %d\n", counter)
	cond.Signal() // cond.Signal() wakes up sleeping goroutines one by one

	mutex.Unlock()
}
