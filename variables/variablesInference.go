package main

import "fmt"

func main() {
	a := 42
	b := 3.14159
	c := 0.123 + 1.345i
	d := "Hello World It's golang"

	fmt.Printf("[a] is the type of %T\n", a) // int
	fmt.Printf("[b] is the type of %T\n", b) // float64
	fmt.Printf("[c] is the type of %T\n", c) // complex128
	fmt.Printf("[d] is the type of %T\n", d) // string
}
