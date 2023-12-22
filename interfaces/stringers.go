package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// A Stringer is a type that can describe itself as a string
// The fmt package (any many others) look for this interfface to print values
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Knightchaser", 20}
	b := Person{"SomeCPO", 35}
	fmt.Println(a)
	fmt.Println(b)
}
