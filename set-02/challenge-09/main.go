package main

import (
	"fmt"
)

func main() {
	blocksize := 20
	block := "YELLOW SUBMARINE"

	padded := padFinalBlock(blocksize, []byte(block))
	fmt.Println(padded)
}

func padFinalBlock(blocksize int, block []byte) []byte {
	p := blocksize - len(block)
	for i := 0; i < p; i++ {
		block = append(block, byte(p))
	}
	return block
}
