// sync.Once guarantees that the target function will be executed "only once"
// even though the multiple goroutines are trying to run the target function multiple times
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func hello() {
	fmt.Println("This will be executed only once :)")
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	syncOnce := new(sync.Once)
	goroutineQty := 5

	for n := 0; n < goroutineQty; n++ {
		go func(id int) {
			fmt.Printf("Here is the goroutine %d\n", id)
			syncOnce.Do(hello) // function "hello" will be run only once in the entire program.
		}(n)
	}

	time.Sleep(2 * time.Second)

}
