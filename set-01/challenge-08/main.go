package main

import (
	"crypto/aes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file := "./data.txt"
	ecbLine := detectAESinECB(file)
	fmt.Println(ecbLine)
}

// ==========
// Master function - run this in main
// ==========
func detectAESinECB(f string) string {
	data := readDataFile(f)
	lines := splitFileLines(data)

	reps := 0
	var ecb string

	for i := 0; i < len(lines); i++ {
		str, count := countRepeatingBlocks(lines[i])
		if count >= reps {
			reps = count
			ecb = str
		}
	}
	return ecb
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

func splitFileLines(s string) []string {
	return strings.Split(s, "\n")
}

func countRepeatingBlocks(s string) (string, int) {
	b := make([]string, len(s)/aes.BlockSize)
	sum := int(0)
	for i := 0; i < len(s); i += aes.BlockSize {
		block := s[i : i+aes.BlockSize]
		b = append(b, block)
		for _, val := range b {
			if val == block {
				sum++
			}
		}
	}
	return s, sum
}
