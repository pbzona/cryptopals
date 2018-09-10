package main

import (
	"encoding/hex"
	"fmt"
)

const input = string("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

func main() {
	inp, _ := hex.DecodeString(input)
	r := getReadableChars()

	var highScore = float32(0)
	var key string
	var plaintext string

	for _, char := range r {
		keyByte := char
		ptext := xor(inp, keyByte)

		var score = float32(0)
		for _, letter := range ptext {
			for _, b := range r {
				if letter == b {
					score++
				}
			}
		}

		if score > highScore {
			highScore = score
			key = string(char)
			plaintext = string(ptext)
		}
	}

	fmt.Println("Score:", highScore, "\nKey:", key, "\nPlaintext:", plaintext)
}

func xor(a []byte, b byte) []byte {
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b
	}
	return c
}

func getReadableChars() []byte {
	// From: http://www.asciitable.com/
	r := []byte{}

	// Space
	r = append(r, byte(20))
	// A-Z
	for i := 65; i <= 90; i++ {
		r = append(r, byte(i))
	}
	// a-z
	for i := 97; i <= 122; i++ {
		r = append(r, byte(i))
	}

	return r
}
