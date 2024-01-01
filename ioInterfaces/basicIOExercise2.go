package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"ioInterfaces/basicIOExercise2.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%d, %f, %s", 1, 2.71828, "fmt.Fprintf()\n") // formatted print to the file
	writer.Flush()                                                   // clear the buffer, by saving buffer content to the file
}
