package main

import (
	"fmt"
	"testing"
)

func TestImplementRepeatingKeyXOR(t *testing.T) {
	const input = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	const key = "ICE"
	const expected = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	output := implementRepeatingKeyXOR(input, key)

	fmt.Println("Expected:", expected)
	fmt.Println("Got:     ", output)

	if output != expected {
		t.Error("Nope!")
	}
}
