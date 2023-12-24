package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type WeirdObject struct {
	WeirdLength float64
}

func printArea(s interface{}) {
	// Type switch to determine the underlying type
	switch v := s.(type) {
	case Rectangle:
		fmt.Printf("Type: Rectangle, Area: %f\n", v.Area())
	case Circle:
		fmt.Printf("Type: Circle, Area: %f\n", v.Area())
	default:
		fmt.Printf("Type: %T, Unexpected object type\n", v)
	}
}

func main() {
	// Create instances of Rectangle and Circle
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}
	weirdObject := WeirdObject{WeirdLength: 10}

	printArea(rectangle)
	printArea(circle)
	printArea(weirdObject)
}
