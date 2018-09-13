package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	const src = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	const key = "ICE"

	res := implementRepeatingKeyXOR(src, key)
	fmt.Println(res)
}

// =========
// Master function - this runs in main
// =========
func implementRepeatingKeyXOR(s string, k string) string {
	// convert src and key to byte slices
	src := []byte(s)
	key := []byte(k)

	// generate a repeating key
	rep := getRepeatingKey(key, src)

	// XOR against src
	out := []byte{}
	for i, letter := range src {
		out = append(out, xor(letter, rep[i]))
	}

	// convert back to hex
	result := hex.EncodeToString(out)

	return result
}

// =========
// Component functions - these run in master
// =========

// Helper - bytewise XOR
func xor(a byte, b byte) byte {
	return a ^ b
}

// Generate a repeating key
func getRepeatingKey(k []byte, ptx []byte) []byte {
	return bytes.Repeat(k, len(ptx))
}
