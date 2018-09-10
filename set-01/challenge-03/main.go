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
	fq := make(map[string]int)

	fq["e"] = 26
	fq["t"] = 25
	fq["a"] = 24
	fq["o"] = 23
	fq["i"] = 22
	fq["n"] = 21
	fq["s"] = 20
	fq["r"] = 19
	fq["h"] = 18
	fq["d"] = 17
	fq["l"] = 16
	fq["u"] = 15
	fq["c"] = 14
	fq["m"] = 13
	fq["f"] = 12
	fq["y"] = 11
	fq["w"] = 10
	fq["g"] = 9
	fq["p"] = 8
	fq["b"] = 7
	fq["v"] = 6
	fq["k"] = 5
	fq["x"] = 4
	fq["q"] = 3
	fq["j"] = 2
	fq["z"] = 1

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
