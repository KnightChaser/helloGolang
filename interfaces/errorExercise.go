package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt should return a non-nil error value when given a negative number,
// as it doesn't support complex numbers.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	var testingnumber = [...]float64{2, -2, 7, -7}
	for i := 0; i < len(testingnumber); i++ {
		numberOut, errorOut := Sqrt(testingnumber[i])
		fmt.Printf("(%v, %v)\n", numberOut, errorOut)
	}
}
