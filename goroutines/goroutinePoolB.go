package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct {
	tag    string // pooling tag
	buffer []int  // data slice
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Allocate a new pool
	// If there is an object in the pool, the existing object will be returned instead
	pool := sync.Pool{
		// A definition of Get() (pool.Get())
		New: func() interface{} {
			data := new(Data)
			data.tag = "new"
			data.buffer = make([]int, 10)
			return data // allocate a new memory
		},
	}

	// Create 10 goroutines
	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data) // Get object and do a type assertion (mandatory)
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // fill random values in the given slice
			}
			fmt.Printf("Goroutine #1, Data => %v\n", data)
			data.tag = "used" // set the object tag as "used"
			pool.Put(data)    // save object to the pool
		}()
	}

	// Create 10 goroutines
	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n // fill even numbers in the given slice
				n += 2
			}
			fmt.Printf("Goroutine #2, Data => %v\n", data)
			data.tag = "used" // set the object tag as "used"
			pool.Put(data)    // save object to the pool
		}()
	}

	fmt.Scanln()

}
