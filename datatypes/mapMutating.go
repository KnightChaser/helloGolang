package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer") // delete data which key is "Answer" from the map "m"
	fmt.Println("The value: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Printf("The value: %d, Present?: %v\n", v, ok)

}
