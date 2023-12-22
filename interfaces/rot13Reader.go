package main

import (
	"io"
	"os"
	"strings"
)

// Let's modify the string in some way
type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rr.r.Read(p)
	for i := 0; i < n; i++ {
		b := p[i]
		// decoding mechanism
		if 'A' <= b && b <= 'Z' {
			p[i] = (b-'A'+13)%26 + 'A'
		} else if 'a' <= b && b <= 'z' {
			p[i] = (b-'a'+13)%26 + 'a'
		} else {
			p[i] = b
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // Realy &r to standard output stream
}
