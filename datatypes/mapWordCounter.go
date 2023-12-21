package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	list := make(map[string]int)
	for _, word := range words {
		list[word] += 1
	}
	return list
}

func main() {
	a := "Let's implement wordcount, it will return the count of splitted string s and the count of the frequency of each element."
	wordCountList := WordCount(a)
	for key, value := range wordCountList {
		fmt.Printf("word[%s] => %d times occurred\n", key, value)
	}
}
