package main

import (
	"strings"
)

func charList(input string) string {
	var char string
	if strings.Contains(input, "L") {
		char += L
	}
	if strings.Contains(input, "U") {
		char += U
	}
	if strings.Contains(input, "N") {
		char += N
	}
	if strings.Contains(input, "S") {
		char += S
	}
	return char
}

func checkAdd(input string) string {
	const swap = N + S
	var char string
	for i := 0; i < len(swap); i++ {
		if strings.Contains(input, swap[i]) {
			char += swap
		}
	}

	return char
}

const ( // change letters later??
	// Lowercase letters - 26
	L = "abcdefghijklmnopqrstuvwxyz"
	// Uppercase letters - 26
	U = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numbers - 10
	N = "0123456789"
	// Symbols - 30
	S = "!@#$%^&*()-_=+[]{}|;:,.<>?/~"
)
