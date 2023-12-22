package main

import (
	"fmt"
	"os"
)

func main() {

	// 1. Trying to open the invalid file that doesn't exist actually
	openFile("Invalid.txt")

	// 5. Because the program is not panic now, this will be also done.
	fmt.Println("Done")
}

func openFile(fn string) {

	// 3. Deferred functions will be executed as openFile() closes
	defer func() {
		// 4. Initiate recover() session, panic status transits to normal status
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		// 2. Panic happenes due to an I/O error
		panic(err)
	}

	defer f.Close()
}
