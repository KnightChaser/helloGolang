package main

import (
	"fmt"
	"time"
)

// Create a new type error struct
type MyError struct {
	When time.Time
	What string
}

// Similar to fmt.Stringer (overrriding in C++, actually)
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	fmt.Println("Seems to be working for now? Let's see about that")
	return &MyError{
		time.Now(),
		"Anyway it didn't work so well >.0",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
