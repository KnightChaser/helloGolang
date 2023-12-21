// Defer delays a specific function's execution until the surrounding function finishes working
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World") // It prints later because the surrounding function must be executed in prior
	fmt.Println("Hello")       // It prints first
}
