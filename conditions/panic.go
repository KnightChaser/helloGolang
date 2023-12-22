package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("I'm trying to open a file(deferred print)")

	// trying to open an invalid file, causing paninc
	openFile("Invalid.txt")

	// because the function got a panic, the subsequence instructions won't be executed at all
	defer fmt.Println("Done (deferred print)")
	println("Done")
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
