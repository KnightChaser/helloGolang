// In Go, Golang doesn't run the following instructions after executing only the first exact match
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on...")
	// runtime.GOOS returns the current operating system that this golang runs on as a string
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("Something else(currently, %s)\n", os)
	}
}
