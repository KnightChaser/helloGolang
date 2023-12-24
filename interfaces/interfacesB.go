package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

// Define rectangle
type Rect struct {
	width  float64
	height float64
}

// Define circle
type Circle struct {
	radius float64
}

// Implement Shape interface for type of Rect(Rectangle)
func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width + r.height) }

// Implement Shape interface for type of Circle(Circle)
func (c Circle) area() float64      { return math.Pi * math.Pow(c.radius, 2) }
func (c Circle) perimeter() float64 { return 2 * math.Pi * c.radius }

// The func showShapeInformation() can hold multiple shape type objects
func showShapeInformation(shapes ...Shape) {
	for number, shape := range shapes {
		// Call shape object
		areaValue := shape.area()
		perimeterValue := shape.perimeter()
		fmt.Printf("Object #%v Area: %v\n", number, areaValue)
		fmt.Printf("Object #%v Perimeter: %v\n", number, perimeterValue)
	}
}

func main() {
	exampleRectangle := Rect{10, 20}
	exampleCircle := Circle{12}
	showShapeInformation(exampleRectangle, exampleCircle)
}
