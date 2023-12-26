package main

import (
	"fmt"
)

// Via interface, define a type set for generic functions
type ComparableNumberIntegers interface {
	int | int16 | int32 | int64
}

type ComparableNumberFloatings interface {
	float32 | float64
}

// merge the two prior interfaces for generic constraints
type ComparableNumberTypes interface {
	ComparableNumberIntegers | ComparableNumberFloatings
}

func max[T ComparableNumberTypes](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	var (
		a int     = 1
		b int     = 129
		c int16   = 3333
		d int16   = 8181
		e float32 = 3.14159
		f float32 = 20.83
	)
	fmt.Println(max(a, b))
	fmt.Println(max(c, d))
	fmt.Println(max(e, f))
}
