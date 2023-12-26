package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true // After certain tasks, sends "true" via the boolean channel done
	}()

	fmt.Println("?")
	<-done // Wait until the upper goroutine finishes working
	fmt.Println("Done!")
}
