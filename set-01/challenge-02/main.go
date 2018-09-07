package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	p := "1c0111001f010100061a024b53535009181c"
	k := "686974207468652062756c6c277320657965"

	fmt.Println(fixedXor(p, k))
}

// Master function
func fixedXor(plaintext string, key string) string {
	plaintextBytes := hexToBytes(plaintext)
	keyBytes := hexToBytes(key)

	return bytesToHex(xor(plaintextBytes, keyBytes))
}

// Component functions
func hexToBytes(input string) []byte {
	b, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func xor(a []byte, b []byte) []byte {
	var c []byte
	for i := range a {
		c = append(c, a[i]^b[i])
	}
	return c
}

func bytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}
