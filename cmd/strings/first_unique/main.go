package main

import (
	"fmt"
	"golang-exercises/internal/strings/unique"
)

func main() {
	input := "AAAABBBBBYKKK"
	result := unique.FirstUniqChar(input)

	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Result, %d\n", result)
}
