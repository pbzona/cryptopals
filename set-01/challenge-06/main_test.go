package main

import (
	"fmt"
	"testing"
)

// This test given in the problem statement
func TestHammingDistance(t *testing.T) {
	first := []byte("this is a test")
	second := []byte("wokka wokka!!!")
	dist := hammingDistance(first, second)

	if dist != 37 {
		fmt.Println("Calculated distance:", dist)
		t.Error("Nope!")
	}
}
