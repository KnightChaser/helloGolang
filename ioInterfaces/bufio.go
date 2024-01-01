package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile(
		"ioInterfaces/hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		return
	}
	defer file.Close() // Don't forget to close the file object before main() ends

	writer := bufio.NewWriter(file)                // Create a new io.Writer instance to the file obj
	writer.WriteString("Hello Hello Happiness!\n") // Write a string
	writer.Flush()                                 // clear buffer by saving every content in the buffer to the file

	reader := bufio.NewReader(file)
	fileInfo, _ := file.Stat()
	b := make([]byte, fileInfo.Size()) // Create a byte slice with file size amount

	file.Seek(0, os.SEEK_SET) // starting from the first page of file
	reader.Read(b)            // Read file and store the result to the buffer
	fmt.Printf("File read: %v\n", string(b))

}
