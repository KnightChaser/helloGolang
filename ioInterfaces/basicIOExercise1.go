package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Hello, World"
	reader := strings.NewReader(s)
	var string1, string2 string
	// fmt.Fscanf() function can use file, io.Reader, std I/O instances
	inputQty, _ := fmt.Fscanf(reader, "%s %s", &string1, &string2)

	fmt.Printf("Input QTY: %d\n", inputQty)
	fmt.Println(string1)
	fmt.Println(string2)

}
