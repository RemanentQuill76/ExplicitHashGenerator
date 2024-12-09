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
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for {
		fmt.Print("Please enter the filepath/name of the file you would like to read from: ")
		if scanner.Scan() {
			filePath := scanner.Text()
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				fmt.Println("File not found. Please try again.")
				continue
			}
			lines = read_file(filePath)
			break
		} else {
			fmt.Println("Failed to read input")
			return
		}
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
