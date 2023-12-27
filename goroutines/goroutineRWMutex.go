// Read/Write Mutex(RW Mutex) can lock reads and writes separately
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	var data int = 0
	var rwMutex = new(sync.RWMutex) // Create a read/write MUTEX(Mutual exclusion)
	// Such Read/Write Mutex is very useful to block reading/writing previous data
	// when the important data IO task is being processed. It also prevents the data unexpectedly
	// changed during reading. Especially, RWMutex is more advantageous when reading is often than writing.

	// A goroutine that writes data
	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // Lock mutex
			data += 1
			fmt.Printf("Write: %v\n", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock() // Unlock mutex
		}
	}()

	// A goroutine that reads data
	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // Lock read mutex
			fmt.Printf("read(goroutine #1): %v\n", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock() // Unlock read mutex
		}
	}()

	// A goroutine that reads data
	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // Lock read mutex
			fmt.Printf("read(goroutine #2): %v\n", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock() // Unlock read mutex
		}
	}()

	// Run the goroutine for 10 seconds
	time.Sleep(10 * time.Second)

}
