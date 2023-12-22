// Readers (io.Reader)
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Trying to use Golang Readers!") // Create strings.Reader
	b := make([]byte, 8)                                    // Consumes 8 bytes at a time
	for {
		n, err := r.Read(b)
		fmt.Printf("n => %v | err => %v | b => %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n\n", b[:n])
		if err == io.EOF {
			fmt.Printf("(Reading got finished)\n")
			break
		}
	}
}
