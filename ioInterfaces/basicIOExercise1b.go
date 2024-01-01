package main

import (
	"fmt"
	"os"
)

func main() {
	// Open thSe file for reading
	file, err := os.Open("ioInterfaces/basicIOExercise1b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var string1, string2 string
	// fmt.Fscanf() function can use file, io.Reader, std I/O instances
	inputQty, err := fmt.Fscanf(file, "%s %s", &string1, &string2)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Printf("Input QTY: %d\n", inputQty)
	fmt.Println(string1)
	fmt.Println(string2)
}
