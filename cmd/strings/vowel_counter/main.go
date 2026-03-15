package main

import (
	"fmt"
	"golang-exercises/internal/strings/counter"
)

func main() {
	input := "Golang Exercises"
	result := counter.VowelCounter(input)

	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Result, Total Vowel: %d\n", result)
}
