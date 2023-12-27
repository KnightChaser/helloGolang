package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("This is goroutine")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("hello hello")

	wait.Wait()
}
