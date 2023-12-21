package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe        bool       = false
	MaxInt      uint64     = 1<<64 - 1            // 2^64-1
	z           complex128 = cmplx.Sqrt(-5 + 12i) // complex number (in mathematics)
	voidInt     int
	voidBoolean bool
	voidString  string
)

func main() {
	fmt.Printf("ToBe: %T(%v)\n", ToBe, ToBe)
	fmt.Printf("MaxInt: %T(%v)\n", MaxInt, MaxInt)
	fmt.Printf("z: %T(%v)\n", z, z)

	// undeclared variables will be set as zero value as a default
	fmt.Printf("voidInt: %T(%v)\n", voidInt, voidInt)             // 0 as default
	fmt.Printf("voidBoolean: %T(%v)\n", voidBoolean, voidBoolean) // false as default
	fmt.Printf("voidString: %T(%v)\n", voidString, voidString)    // "" as default
}
