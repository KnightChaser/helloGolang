// If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.

package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// Nil receiver
// Other programming languages would trigger nullPointerException at this moment
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	} else {
		fmt.Println(t.S)
	}
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i) // (<nil>, *main.T)
	i.M()       // <nil>

	i = &T{"Hello Golang interfaces"}
	describe(i) // (&{Hello Golang interfaces}, *main.T)
	i.M()       // "Hello Golang interfaces"
}
