package main

import (
	"fmt"
)

func main() {
	NATOPhoneticAlphabet := map[string]string{
		"A": "Alpha",
		"B": "Bravo",
		"C": "Charlie",
		"D": "Delta",
		"E": "Echo",
		"F": "Foxtrot",
		"G": "Golf",
	}

	for key, value := range NATOPhoneticAlphabet {
		fmt.Printf("%s => %s\n", key, value)
	}
}
