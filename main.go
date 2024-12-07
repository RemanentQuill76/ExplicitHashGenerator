package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Print("Welcome to the password generator")
	time.Sleep(1 / 2 * time.Second)
	fmt.Print("Please enter the filepath/name of the file you would like to read from: ")
	// fmt.Print("will you be using cuda today? (y/n)")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	if scanner.Scan() {
		lines = read_file(scanner.Text())
	} else {
		fmt.Println("Failed to read input")
		return
	}
	fmt.Printf("the initial password count is: %d\n", len(lines))
	time.Sleep(1 / 2 * time.Second)
	fmt.Println("Enter permutation depth (1-5):")
	scanner.Scan()
	depthStr := scanner.Text()
	depth, err := strconv.Atoi(depthStr)
	if err != nil || depth < 1 || depth > 5 {
		fmt.Println("Invalid depth. Please enter a number between 1 and 5.")
		return
	}
	fmt.Println("permutation depth is: ", depth)

	// if err:= cu.0); err != nil {
	//   panic(err)
	// }
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

func read_file(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
