package main

import (
	"fmt"
	"math"
)

type Shape interface {
	OuterArea() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) OuterArea() float64 {
	return 6 * math.Pow(c.side, 2)
}

func (c Cube) Volume() float64 {
	return math.Pow(c.side, 3)
}

func main() {
	cube := Cube{10}
	var cubeShapeObj Shape = cube
	var cubeVolumeObj Object = cube
	fmt.Printf("OuterArea of (%T)%v => %v\n", cubeShapeObj, cubeShapeObj, cubeShapeObj.OuterArea())
	fmt.Printf("Volume    of (%T)%v => %v\n", cubeVolumeObj, cubeVolumeObj, cubeVolumeObj.Volume())
}
