package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := "Hello Golang io.Reader()\n"
	r := strings.NewReader(s)

	// io.Copy() is kinda swiss knife of IO redirection
	// copy the data of io.NewReader() instance to stdout(console)
	io.Copy(os.Stdout, r)
}
