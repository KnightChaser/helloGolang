package main

import (
	"fmt"
)

func max[T int | int16 | int32 | int64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	} else {
		return b
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
