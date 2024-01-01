package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile(
		"ioInterfaces/stringToFile.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		return
	}
	defer file.Close()

	targetString := "The string now gonna get into the file\n"
	reader := strings.NewReader(targetString) // Convert the string into the NewReader form
	writer := bufio.NewWriter(file)           // An io.NewWriter() instance grabs the file object
	writer.ReadFrom(reader)                   // Read from the io.NewReader() instance holding string and save it to the writer buffer
	writer.Flush()                            // Flush io.NewWriter() by saving the data to the file

	// Alternative
	reader = strings.NewReader(targetString) // Read from the
	io.Copy(writer, reader)                  // Copy io.NewReader() instance data to the Writer instance
	writer.Flush()                           // Save to the Writer instance
}
