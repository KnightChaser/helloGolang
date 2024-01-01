package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	exampleCredential := `
	{
		"username": "kn1ght",
		"password": "OwOCredential$$1212",
		"hostname": "klojure.kr-AZ6.owodb.org",
		"port"	  : 12345
	}
	`

	type credentialInformation struct {
		Username string
		Password string
		Hostname string
		Port     uint
	}

	var credential credentialInformation

	json.Unmarshal([]byte(exampleCredential), &credential)
	fmt.Printf("%s@%s:%d (PW: %s)\n", credential.Username, credential.Hostname, credential.Port, credential.Password)
	fmt.Printf("Raw JSON: %v\n", credential)

}
