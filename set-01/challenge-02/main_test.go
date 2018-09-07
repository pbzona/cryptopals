package main

import (
	"testing"
)

func TestConversion(t *testing.T) {
	plaintext := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	expect := "746865206b696420646f6e277420706c6179"

	if fixedXor(plaintext, key) != expect {
		t.Error("Nope!")
	}
}
