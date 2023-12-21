// By traversing slices, it's able to copy the whole data in array like obj
package main

import (
	"fmt"
)

func main() {
	slice1 := []int{64, 128, 256, 512, 1024}
	slice2 := make([]int, len(slice1))
	fmt.Println(slice1, slice2)
	for i, v := range slice1 {
		fmt.Printf("slice1[%d] => %d\n", i, v)
		slice2[i] = v
	}
	fmt.Println(slice1, slice2)
}
