package main

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	const file = "./data.txt"
	const key = "YELLOW SUBMARINE"
	iv := make([]byte, aes.BlockSize)

	fmt.Println(implementCBCMode(file, key, iv))
}

// ==========
// Master function - run this in main
// ==========
func implementCBCMode(in string, key string, iv []byte) string {
	data := readDataFile(in)
	c, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	if len(c) < cipher.BlockSize() {
		err = errors.New("Ciphertext block size too short")
		panic(err)
	}

	res := make([]string, len(c))
	prevBlock := make([]byte, cipher.BlockSize())
	prevBlock = iv
	for i := 0; i < len(c); i += cipher.BlockSize() {
		currentBlock := c[i : i+cipher.BlockSize()]
		t := make([]byte, len(currentBlock))
		copy(t, currentBlock)
		cipher.Decrypt(currentBlock, currentBlock)
		ptbytes := xor(currentBlock, prevBlock)
		res = append(res, string(ptbytes))
		prevBlock = t
	}

	return strings.Join(res, "")
}

// ==========
// Component functions - run these in master
// ==========
func readDataFile(f string) string {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func xor(src []byte, key []byte) []byte {
	res := []byte{}
	for i, b := range src {
		res = append(res, (b ^ key[i]))
	}
	return res
}
