package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	x := 12 // x position in the given image
	y := 12 // y position in the given image
	fmt.Printf("image boundary => %v\n", m.Bounds())
	fmt.Printf("image data at (%d, %d) => %v\n", x, y, m.At(x, y))
}
