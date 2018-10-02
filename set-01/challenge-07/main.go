package main

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	file := "./data.txt"
	key := []byte("YELLOW SUBMARINE")

	fmt.Println(AESInECBMode(file, key))
}

// ==========
// Master function - run this in main
// ==========
func AESInECBMode(f string, k []byte) string {
	data := readDataFile(f)
	return decryptCiphertext(data, k)
}

// ==========
// Component functions - run these in master
// ==========

// Get data from the file on disk
func readDataFile(f string) string {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return string(file)
}

// Decrypt AES-128
func decryptCiphertext(src string, key []byte) string {
	c, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(c) < aes.BlockSize {
		err = errors.New("Ciphertext block size too short")
	}

	for i := 0; i < len(c); i += 16 {
		block.Decrypt(c[i:i+16], c[i:i+16])
	}
	return string(c)
}
