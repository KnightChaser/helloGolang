package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	message := "Hello, Golang SHA"
	hashDigest1 := sha512.Sum512([]byte(message)) // calculate hash digest
	fmt.Printf("SHA512: %x\n", hashDigest1)

	hashDigestInstance := sha512.New()
	hashDigestInstance.Write([]byte("Hello, "))
	hashDigestInstance.Write([]byte("Golang SHA"))
	hashDigest2 := hashDigestInstance.Sum(nil)
	fmt.Printf("SHA512: %x\n", hashDigest2)
}
