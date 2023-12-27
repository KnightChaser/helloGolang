package main

import (
	"fmt"
	"sync"
)

// Object to be stored in the pool
type MyObject struct {
	Value int
}

func main() {
	// Create a new sync.Pool
	pool := &sync.Pool{
		New: func() interface{} {
			// This function is called when a new object is needed
			return &MyObject{}
		},
	}

	// Acquire an object from the pool
	obj1 := pool.Get().(*MyObject)
	obj1.Value = 42
	fmt.Println("Object 1:", obj1)

	// Release the object back to the pool
	pool.Put(obj1)

	// Acquire another object from the pool
	obj2 := pool.Get().(*MyObject)
	fmt.Println("Object 2:", obj2) // (*)the object is reused. It has the value of obj1.Value (42)

	// Note that the Value field is not reset because the object was reused from the pool
	// If the New function is specified, it will be called when a new object is needed
	// Otherwise, the zero value for the type will be used

	// Release the object back to the pool
	pool.Put(obj2)
}
