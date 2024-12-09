package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

struct filter {
	// 0 values mean not applicable 
	// 1 values mean applicable
	name string
	lenTop int
	lenBottom int
	contains []string
	mathRule
}


func importFilterFile()	{
	fmt.Print("What file do you want to filter with? ")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	if scanner.Scan() {
		lines = read_file(scanner.Text())
	} else {
		fmt.Println("Failed to read input")
		return
	}
}

func applyFilter(filter[]string, lines[]string)	{
	fmt.Println("Applying filter")
	if len(filter) == 0 {
		fmt.Println("Filter is empty")
		return
	}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(filter); j++ {
			if lines[i] == filter[j] {
				fmt.Println(lines[i])
			}
		}
	}

}