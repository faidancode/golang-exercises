package main

import (
	"fmt"
	"golang-exercises/internal/strings/reverser"
)

func main() {
	input := "Hello World"
	result := reverser.Reverse(input)

	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Result: %s\n", result)
}
