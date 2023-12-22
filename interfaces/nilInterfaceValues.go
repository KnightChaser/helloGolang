package main

import (
	"fmt"
)

// A nil interface value holds neither value nor concrete types
type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	// nothing is written, so nil will be
	describe(i)
	i.M()
}
