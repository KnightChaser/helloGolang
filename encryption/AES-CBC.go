package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func encryptAESCBC(block cipher.Block, plaintext []byte) []byte {

	// padding for block size
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}

	// Create ciphertext area with IV(Initial Vector)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	initialVector := ciphertext[:aes.BlockSize] // grab IV area by sub slice
	if _, error := io.ReadFull(rand.Reader, initialVector); error != nil {
		log.Fatal(error)
		return nil
	}

	// Proceed encryption (Lead IV behind)
	mode := cipher.NewCBCEncrypter(block, initialVector)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

func decryptAESCBC(block cipher.Block, ciphertext []byte) []byte {

	// Check padding and block size validity
	if len(ciphertext)%aes.BlockSize != 0 {
		log.Fatal("The encrypted message should be k times of block size in AES cryptosystem.\n")
		return nil
	}

	// Create plaintext area
	initialVector := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:] // Detach IV from ciphertext chunk
	plaintext := make([]byte, len(ciphertext))

	// Process decryption
	mode := cipher.NewCBCDecrypter(block, initialVector)
	mode.CryptBlocks(plaintext, ciphertext)

	return plaintext
}

func main() {
	key := "KEYSIZEMUSTBE16B"
	message := "AES is an abbreviation of \"Advanced Encryption Standard\", which is one of the most popular and strong cryptosystem!\n"
	block, error := aes.NewCipher([]byte(key))
	if error != nil {
		log.Fatal(error)
		return
	}

	ciphertext := encryptAESCBC(block, []byte(message))
	fmt.Printf("Encrypted: %x\n", ciphertext)

	plaintext := decryptAESCBC(block, ciphertext)
	fmt.Printf("Decrypted: %s\n", string(plaintext))
}
