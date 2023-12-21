package main

import (
	"fmt"
	"math"
)

func main() {
	pow := []int{}
	for i := 0; i <= 10; i++ {
		result := int(math.Pow(2, float64(i)))
		pow = append(pow, result)
	}

	for i, v := range pow {
		fmt.Printf("array[%d] => %d\n", i, v)
	}
}
