package main

import "fmt"

func main() {
	a, b := 1024, 2048
	str := "hello Golang function closure"

	// anonymous function
	result := func() int {
		// access to the variables in the main function becausue
		// it's an anonymous function, thus it's also a function closusre
		return a + b
	}()

	// anonymous function
	func() {
		fmt.Println(str)
	}()

	fmt.Println(result)
}
