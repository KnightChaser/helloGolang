// Assertion check is the validation of the data type for some specific variables
// Type assertion provides access to an interface value's underlying concrete value.
package main

import (
	"fmt"
)

func main() {

	var i interface{} = "hello world"

	s := i.(string)
	fmt.Println(s) // "hello world"

	s, ok := i.(string)
	fmt.Println(s, ok) // "hello world" true

	// if you use two variables to get type assertions, it won't panic if the type assertion is not satisfied.
	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panicking because it is string, not float64
	fmt.Println(f)

}
