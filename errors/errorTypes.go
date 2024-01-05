package main

import (
	"fmt"
	"log"
	"time"
)

type HelloOneCustomError struct {
	time  time.Time // the time when the error triggered
	value int       // the value that triggered the error
}

func (e HelloOneCustomError) Error() string {
	return fmt.Sprintf("%v: %d is not 1, triggering helloOneCustomError\n", e.time, e.value)
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintf("Hello %d", n)
		return s, nil
	}
	return "", HelloOneCustomError{time.Now(), n}
}

func main() {

	// normal case
	s, err := helloOne(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	// abnormal(error) case
	s, err = helloOne(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	fmt.Println("The code does not executed since log.Fatal()")

}
