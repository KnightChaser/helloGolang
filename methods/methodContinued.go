package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// Or... you can use type float64 instead.
// Then, you might have to convert MyFloat to float64(). (*type casating)
// The receiver function should be having the same type. (e.g. Int won't be accpeted by this receiver function)
func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return -f
	} else {
		return f
	}
}

func main() {
	f := MyFloat(-math.Sqrt2) // math constant of sqrt(2) * (-1)
	fmt.Printf("f.Abs() => %v\n", f.Abs())
}
