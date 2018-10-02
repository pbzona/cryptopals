package main

import (
	"fmt"
)

func main() {
	file := "./data.txt"
	breakRepeatingKeyXor(file)
}

// ==========
// Master function - run this in main
// ==========
func breakRepeatingKeyXor(filename string) {
	data := readDataFile(filename)
	keyLen := guessKeySize([]byte(data))[0].keySize
	blocks := transposeBlocks(getBlocks(data, keyLen))

	key := []byte{}
	for _, block := range blocks {
		key = append(key, []byte(singleByteXor(block))[0])
	}

	plaintext := implementRepeatingKeyXOR(data, key)
	fmt.Println(plaintext)
}
