package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
)

func main() {
	// DO NOT use this salt value; generate your own random salt. 8 bytes is
	// a good length.
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

	dk, err := scrypt.Key([]byte("some password"), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(1<<15)
	fmt.Println(base64.StdEncoding.EncodeToString(dk))
}
