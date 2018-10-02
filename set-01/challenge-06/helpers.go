package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"math"
	"sort"
)

// ==========
// Component functions - run these in master
// ==========

// Calculate the Hamming distance between two byte slices
func hammingDistance(txt1, txt2 []byte) int {
	switch bytes.Compare(txt1, txt2) {
	case 0: // txt1 == txt2
	case 1: // txt1 > txt2
		temp := make([]byte, len(txt1))
		copy(temp, txt2)
		txt2 = temp
	case -1: // txt1 < txt2
		temp := make([]byte, len(txt2))
		copy(temp, txt1)
		txt1 = temp
	}

	if len(txt1) != len(txt2) {
		panic("Undefined for sequences of unequal length")
	}

	count := 0
	for idx, b1 := range txt1 {
		b2 := txt2[idx]
		xor := b1 ^ b2
		// bit count (number of 1s)
		// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetNaive
		// repeat shifting from left to right (divide by 2)
		// until all bits are zero
		for x := xor; x > 0; x >>= 1 {
			// check if lowest bit is 1
			if int(x&1) == 1 {
				count++
			}
		}
	}
	if count == 0 {
		return 1
	}
	return int(count)
}

// Get data from the file on disk
func readDataFile(f string) []byte {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	b, err := base64.StdEncoding.DecodeString(string(file))
	return b
}

// Guess the size of the repeating key based on Hamming distance
//
// Returning a slice of distances for debugging mostly - we really only
// need the distance of the best result
//
type distance struct {
	keySize     int
	hammingDist float32
}

func guessKeySize(data []byte) []distance {
	// Try key sizes in a given range
	keyMin := 2
	keyMax := 40

	dist := make([]distance, keyMax-keyMin)
	for i := 0; i < keyMax-keyMin; i++ {
		k := i + keyMin
		// Get first N sequences of the slice
		// (More sequences yields more accurate avg distance)
		b1 := data[:k]
		b2 := data[k : k*2]
		b3 := data[k*2 : k*3]
		b4 := data[k*3 : k*4]
		b5 := data[k*4 : k*5]
		b6 := data[k*5 : k*6]
		b7 := data[k*6 : k*7]
		b8 := data[k*7 : k*8]
		b9 := data[k*8 : k*9]
		b10 := data[k*9 : k*10]

		// Get average Hamming distance and normalize
		d := float32(hammingDistance(b1, b2)+hammingDistance(b2, b3)+hammingDistance(b3, b4)+hammingDistance(b4, b5)+hammingDistance(b5, b6)+hammingDistance(b6, b7)+hammingDistance(b7, b8)+hammingDistance(b8, b9)+hammingDistance(b9, b10)) / float32(9*k)
		dist[i] = distance{
			keySize:     k,
			hammingDist: d,
		}
	}

	sort.Slice(dist[:], func(i, j int) bool {
		return dist[i].hammingDist < dist[j].hammingDist
	})

	return dist
}

// Break the ciphertext into blocks of an arbitrary size
func getBlocks(c []byte, size int) [][]byte {
	n := int(math.Ceil(float64(len(c)) / float64(size)))
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		if len(c[i*size:]) >= size {
			b[i] = c[i*size : (i+1)*size]
		} else {
			b[i] = c[i*size:]
		}
	}
	return b
}

// Transpose blocks of arbitrary key size
//
// e.g. First transposed block from first elements of prev block,
// second transposed block from second elements of prev block
//
func transposeBlocks(b [][]byte) [][]byte {
	t := make([][]byte, len(b[0]))
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			t[j] = append(t[j], b[i][j])
		}
	}
	return t
}

// =========
// Single byte XOR functions
// =========
func singleByteXor(input []byte) string {
	chars := getAllASCIIChars()
	fq := createFrequencyTable()

	var highScore = float32(0)
	var key string

	for _, char := range chars {
		keyByte := char
		ptext := xor(input, keyByte)
		score := float32(0)
		for _, letter := range ptext {
			for _, b := range chars {
				if letter == b {
					score += fq[b]
				}
			}
		}
		if score >= highScore {
			highScore = score
			key = string(char)
		}
	}
	return key
}

func xor(a []byte, b byte) []byte {
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b
	}
	return c
}

func getAllASCIIChars() []byte {
	a := []byte{}
	for i := 0; i <= 255; i++ {
		a = append(a, byte(i))
	}
	return a
}

func createFrequencyTable() map[byte]float32 {
	var frequencyTable = map[string]float32{
		"a": 8.167, "b": 1.492, "c": 2.782, "d": 4.253, "e": 13.702,
		"f": 2.228, "g": 2.015, "h": 6.094, "i": 6.966, "j": 0.153,
		"k": 0.772, "l": 4.025, "m": 2.406, "n": 6.749, "o": 7.507,
		"p": 1.929, "q": 0.095, "r": 5.987, "s": 6.327, "t": 9.056,
		"u": 2.758, "v": 0.978, "w": 2.360, "x": 0.150, "y": 1.974,
		"z": 0.074, " ": 17.0,
		"A": 8.167 / 2, "B": 1.492 / 2, "C": 2.782 / 2, "D": 4.253 / 2, "E": 13.702 / 2,
		"F": 2.228 / 2, "G": 2.015 / 2, "H": 6.094 / 2, "I": 6.966 / 2, "J": 0.153 / 2,
		"K": 0.772 / 2, "L": 4.025 / 2, "M": 2.406 / 2, "N": 6.749 / 2, "O": 7.507 / 2,
		"P": 1.929 / 2, "Q": 0.095 / 2, "R": 5.987 / 2, "S": 6.327 / 2, "T": 9.056 / 2,
		"U": 2.758 / 2, "V": 0.978 / 2, "W": 2.360 / 2, "X": 0.150 / 2, "Y": 1.974 / 2,
		"Z": 0.074 / 2,
	}

	table := make(map[byte]float32, len(frequencyTable))
	for char, value := range frequencyTable {
		charByte := []byte(char)[0]
		table[charByte] = value
	}

	return table
}

// =========
// Repeating key XOR functions
// =========
// We can use this to decrypt as well because the XOR reverses itself
// when run against the same input with same key
func implementRepeatingKeyXOR(src []byte, key []byte) string {
	rep := getRepeatingKey(key, src)
	out := []byte{}
	for i, letter := range src {
		out = append(out, singleXor(letter, rep[i]))
	}
	result := string(out)
	return result
}

func getRepeatingKey(k []byte, ptx []byte) []byte {
	return bytes.Repeat(k, len(ptx))
}

func singleXor(a byte, b byte) byte {
	return a ^ b
}
