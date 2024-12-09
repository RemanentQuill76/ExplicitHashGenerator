package cpuSingleThread

import (
	"strings"
	"fmt"
	"time"
	"strconv"
	"bufio"
	"os"
)

func generate1(input string, chars string) []string {
	arr := [5]int{0, 0, 0, 0, 0}
	result := make([]string, len(arr))
	for i, v := range arr {
		result[i] = strconv.Itoa(v)
	}
	return result
}

func regionPermutations (input string) []string {
	var arr []string
	for i := 0; i < len(input); i++ {
		if strings.Contains(input, "L") {
	}
	return arr
}
func generateAll (all []string, chars string) []string {
	var arr []string
	for i := 0; i < len(all); i++ {
		arr = append(arr, generate1(all[i])...)
	}
	return arr
}

func 

func generateCombinations(mask string) []string {
	// Convert mask into character sets
	charSets := []string{}
	for i := 0; i < len(mask); i++ {
		if mask[i] == '?' && i+1 < len(mask) {
			charSets = append(charSets, charSet(mask[i+1]))
			i++ // Skip next character since it's part of the mask
		} else {
			charSets = append(charSets, string(mask[i]))
		}
	}
	return combine(charSets, "")
}

func juggle(arr []string) []string {
	n := len(arr)
	if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i] // For even-size subsets
		}
	} else {
		arr[0], arr[n-1] = arr[n-1], arr[0] // For odd-size subsets
	}
	return arr
}