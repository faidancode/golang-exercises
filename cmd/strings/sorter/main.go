package main

import (
	"fmt"
	"golang-exercises/internal/strings/sorter"
)

func main() {
	input := "acbsddebh"
	result := sorter.Sorter(input)

	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Result: %s\n", result)
}
