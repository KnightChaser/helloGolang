package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulating some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	const numWorkers = 5
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers have completed their tasks")
}
