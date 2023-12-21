package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i // pointer *p holds an address to i
	fmt.Println(*p)
	*p = 21 // the value that the pointer *p changes to 21
	fmt.Printf("%d <---> %d\n", *p, i)

	p = &j
	*p = *p / 37
	fmt.Printf("%d <---> %d\n", *p, j)

	// BTW, golang doesn't support pointer arithmetic
}
