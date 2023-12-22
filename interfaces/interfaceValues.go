package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

// ----------------------

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// ----------------------

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ----------------------

func main() {
	var i I
	i = &T{"Hello"} // this case is string
	describe(i)     // (&{Hello}, *main.T)
	i.M()           // "Hello"

	i = F(math.Pi) // this case is float64(F)
	describe(i)    // (3.1415... main.F)
	i.M()          // 3.1415...
}
