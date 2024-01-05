package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

func main() {
	// Create key set in RSA (Public key is included in the private key object)
	privateKey, error := rsa.GenerateKey(rand.Reader, 2048)
	if error != nil {
		log.Fatal(error)
	}
	publicKey := &privateKey.PublicKey

	message := "RSA, the cryptosystem of factorization."

	ciphertext, error := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	fmt.Printf("Plaintext: %s\n", message)
	// fmt.Printf("key(public): %x\n", publicKey)
	// fmt.Printf("key(private): %x\n", privateKey)
	fmt.Printf("Encrypted: %x\n", ciphertext)

	plaintext, error := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	fmt.Printf("Decrypted: %s\n", string(plaintext))
}
