package main

import (
	"fmt"
	"log"
)

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}
	return "", fmt.Errorf("%d is not 1\n", n)
}

func main() {

	// normal case
	s, err := helloOne(1)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(s)

	// abnormal(error) case
	// The log.Print() won't terminate the process, but the produced logs will be printed
	// on the console after the program termination
	s, err = helloOne(2)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(s)

	fmt.Println("The code does not executed since log.Fatal()")

}
