package main

import (
	"fmt"
	"golang-exercises/internal/numbers"
)

func main() {
	input := 34
	result := numbers.FibonacciChecker(input)

	fmt.Printf("Input : %v\n", input)
	fmt.Printf("Result : %s\n", result)
}
