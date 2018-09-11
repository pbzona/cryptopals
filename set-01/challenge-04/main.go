package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Create a type for the best text match from each line
type decodedLine struct {
	text   string
	key    byte
	score  float32
	cipher string
}

func main() {
	fileName := "./data.txt"

	score, key, plaintext, cipher := detectSingleCharXor(fileName)
	fmt.Println("From line:", cipher, "\n\nScore:", score, "\nKey (ASCII):", key, "\nPlaintext:", plaintext)
}

// ==========
// Master function - put everything together here and call from main
// ==========
func detectSingleCharXor(filepath string) (float32, byte, string, string) {
	strings, _ := readFromFile(filepath)
	lines := splitFileLines(strings)

	allScores := make([]decodedLine, len(lines))
	for i, line := range lines {
		s, k, pt, ct := getBestMatch(line)
		n := decodedLine{
			text:   pt,
			key:    k,
			score:  s,
			cipher: ct,
		}
		allScores[i] = n
	}

	var highScore = float32(0)
	var bestMatch string
	var key byte
	var cipher string

	for _, current := range allScores {
		if current.score > highScore {
			highScore = current.score
			bestMatch = current.text
			key = current.key
			cipher = current.cipher
		}
	}

	return highScore, key, bestMatch, cipher
}

// ==========
// Component functions - used in the "master" function called from main
// ==========

// Read from the file on disk
func readFromFile(filepath string) (string, int) {
	file, _ := os.Open(string(filepath))
	defer file.Close()
	fileinfo, _ := file.Stat()
	size := fileinfo.Size()
	buffer := make([]byte, size)
	bytesread, _ := file.Read(buffer)
	return string(buffer), bytesread
}

// Splits the input on newline
func splitFileLines(s string) []string {
	return strings.Split(s, "\n")
}

// Check against *every* ASCII char as key and return the match with the best score
// See below for function definitions used here
func getBestMatch(input string) (float32, byte, string, string) {
	inp, _ := hex.DecodeString(input)
	a := getAllAsciiChars()
	r := getReadableChars()
	fq := createFrequencyTable()

	var highScore = float32(0)
	var key string
	var plaintext string

	for _, char := range a {
		keyByte := char
		ptext := xor(inp, keyByte)

		var score = float32(0)
		for _, letter := range ptext {
			for _, b := range r {
				if letter == b {
					score += fq[b]
				}
			}
		}

		if score > highScore {
			highScore = score
			key = string(char)
			plaintext = string(ptext)
		}
	}

	return highScore, []byte(key)[0], plaintext, hex.EncodeToString(inp)
}

// Helper - bytewise XOR
func xor(a []byte, b byte) []byte {
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b
	}
	return c
}

// Returns "human readable" ASCII chars
// Note - This is used for the lookup while looping over letters in each XOR'd string; we could use the keys on the frequency table there to save code, but it takes about 5x as long when iterating over the map as opposed to a byteslice
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

// Returns all ASCII chars to use as keys when testing for decryption
func getAllAsciiChars() []byte {
	a := []byte{}
	for i := 0; i <= 255; i++ {
		a = append(a, byte(i))
	}
	return a
}

// Creating the scoring table explicitly from characters so that it can be more easily modified/tweaked
func createFrequencyTable() map[byte]float32 {
	var frequencyTable = map[string]float32{
		"a": 8.167, "b": 1.492, "c": 2.782, "d": 4.253, "e": 13.702,
		"f": 2.228, "g": 2.015, "h": 6.094, "i": 6.966, "j": 0.153,
		"k": 0.772, "l": 4.025, "m": 2.406, "n": 6.749, "o": 7.507,
		"p": 1.929, "q": 0.095, "r": 5.987, "s": 6.327, "t": 9.056,
		"u": 2.758, "v": 0.978, "w": 2.360, "x": 0.150, "y": 1.974,
		"z": 0.074, " ": 16.0, // Value for space character kind of arbitrary, added ~2 to e
		// Divide by 2 for capitals to adjust for real usage in English sentences
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
