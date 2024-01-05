package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

func main() {

	privateKey, error := rsa.GenerateKey(rand.Reader, 2048)
	if error != nil {
		log.Fatal(error)
		return
	}
	publicKey := &privateKey.PublicKey

	// Prepare the data to sign and calculate its hash value
	message := "Hello, RSA signaturing"
	hashMD5 := md5.New()
	hashMD5.Write([]byte(message))
	hashMD5Digest := hashMD5.Sum(nil)

	var signatureHash1 crypto.Hash
	signature, error := rsa.SignPKCS1v15(rand.Reader, privateKey, signatureHash1, hashMD5Digest)
	fmt.Printf("Signed => %x\n", signatureHash1)

	var signatureHash2 crypto.Hash
	error = rsa.VerifyPKCS1v15(publicKey, signatureHash2, hashMD5Digest, signature)
	fmt.Printf("Signed <= %x\n", signatureHash2)

	if error != nil {
		fmt.Println("Failed to verify signature")
	} else {
		fmt.Println("Succeeded to verify signature")
	}
}
