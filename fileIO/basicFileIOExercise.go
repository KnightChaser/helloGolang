package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Open file. Create it if it doesn't exist
	// Use os.O_TRUNC if you wanna delete file content after opening
	file, err := os.OpenFile("./fileIO/filefile.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))

	if err != nil {
		fmt.Printf("An error happened during opening a file => %v", err)
		return
	}

	// Close file at the moment of main()'s termination
	defer file.Close()

	// Write file
	write := bufio.NewWriter(file)
	write.WriteString("Hello Hello Happiness\n")
	write.WriteString("You make me so smile\n")
	write.Flush()

	// Read file again
	read := bufio.NewReader(file)
	fileInformation, _ := file.Stat() // gather file information (metadata)
	fileByteSlice := make([]byte, fileInformation.Size())
	file.Seek(0, os.SEEK_SET)
	read.Read(fileByteSlice)

	fmt.Println(string(fileByteSlice))

}
