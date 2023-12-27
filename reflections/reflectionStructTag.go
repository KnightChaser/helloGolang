package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"name" tag2:"string type"`  // attaching tags to the struct member
	age  int    `tag1:"age"  tag2:"integer type"` // attaching tags to the struct member
}

func main() {

	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name") // struct field name
	fmt.Printf("fieldExist(name) => %v, tag1 => %v, tag2 => %v\n",
		ok, name.Tag.Get("tag1"), name.Tag.Get("tag2")) // retrieving tags

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Printf("fieldExist(age) => %v, tag1 => %v, tag2 => %v\n",
		ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))

}
