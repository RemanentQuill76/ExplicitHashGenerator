package main

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