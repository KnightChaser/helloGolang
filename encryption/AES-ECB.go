package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

func main() {
	// Unified key size and message length as 16 bytes
	key := "KEYSIZEMUSTBE16B"
	message := "EXAMPLEMESSAGEEX"
	block, error := aes.NewCipher([]byte(key))
	if error != nil {
		log.Fatal(error)
		return
	}

	ciphertext := make([]byte, len(message))
	block.Encrypt(ciphertext, []byte(message))
	fmt.Printf("Encrypted: %x\n", ciphertext)

	plaintext := make([]byte, len(message))
	block.Decrypt(plaintext, ciphertext)
	fmt.Printf("Decrypted: %s\n", string(plaintext))
}
