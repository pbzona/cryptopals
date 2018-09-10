package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	breakSingleByteXor(cipher)
}

// Master function
func breakSingleByteXor(cipher string) {
	cipherBytes := hexToBytes(cipher)
	charKeys := getAlphabet()

	// Decodes against each key (based on a basic alphabet of possible keys)
	results := []string{}
	for _, charKey := range charKeys {
		keyByte := []byte(charKey)
		plain := string(xor(cipherBytes, keyByte[0]))
		results = append(results, strings.ToLower(plain))
	}

	// Scores each result against the character frequency table
	scores := []int{}
	for _, result := range results {
		score := scorePlaintext(result, getFrequencyTable())
		scores = append(scores, score)
	}

	// Sort the scores and associated strings
	scoreMap := createScoreMap(scores, results)
	sortScores(scoreMap)
}

// Component functions
func getFrequencyTable() map[string]int {
	// returns a frequency table of ASCII chars based on:
	// http://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
	fq := make(map[string]int)

	fq["e"] = 21912
	fq["t"] = 16587
	fq["a"] = 14810
	fq["o"] = 14003
	fq["i"] = 13318
	fq["n"] = 12666
	fq["s"] = 11450
	fq["r"] = 10977
	fq["h"] = 10795
	fq["d"] = 7874
	fq["l"] = 7253
	fq["u"] = 5246
	fq["c"] = 4943
	fq["m"] = 4761
	fq["f"] = 4200
	fq["y"] = 3853
	fq["w"] = 3819
	fq["g"] = 3693
	fq["p"] = 3316
	fq["b"] = 2715
	fq["v"] = 2019
	fq["k"] = 1257
	fq["x"] = 315
	fq["q"] = 205
	fq["j"] = 188
	fq["z"] = 128

	return fq
}

func getAlphabet() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
}

func hexToBytes(input string) []byte {
	b, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func xor(a []byte, b byte) []byte {
	var c []byte
	for i := range a {
		c = append(c, a[i]^b)
	}
	return c
}

func scorePlaintext(p string, fq map[string]int) int {
	score := 0
	chars := strings.Split(p, "")

	for _, char := range chars {
		if value, hasKey := fq[char]; hasKey {
			score += value
		}
	}

	return score
}

func createScoreMap(scores []int, strings []string) map[int]string {
	scoreMap := make(map[int]string)
	for i, score := range scores {
		scoreMap[score] = strings[i]
	}
	return scoreMap
}

func sortScores(scoreMap map[int]string) {
	scores := []int{}
	for i := range scoreMap {
		scores = append(scores, i)
	}

	// Prints scores in *ascending* order
	sort.Ints(scores)
	for _, score := range scores {
		fmt.Println("Score:", score, "Text:", scoreMap[score])
	}
}
