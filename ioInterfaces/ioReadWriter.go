package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"ioInterfaces/ioReadWriter.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	defer file.Close()

	// io.ReaderWriter() is embedding both io.Reader() and io.Writer()
	reader := bufio.NewReader(file)                               // Create an instance of bufio.NewReader()
	writer := bufio.NewWriter(file)                               // Create an instance of bufio.NewWriter()
	readerWriter := bufio.NewReadWriter(reader, writer)           // Integrate the reader and writer into one instance
	readerWriter.WriteString("Hello, bufio.NewReaderWriter()!\n") // Writer string to the Writer buffer
	readerWriter.Flush()                                          // Flush the buffer to save the content into the buffer

	// Reading...
	file.Seek(0, os.SEEK_SET)
	data, _, _ := readerWriter.ReadLine()
	fmt.Printf("Data: %v\n", string(data))
}
