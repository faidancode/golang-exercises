package main

import (
	"fmt"
	"golang-exercises/internal/array"
)

func main() {
	input := []int{3, 2, 1, 6}
	result := array.ArrayRotation(input)

	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Result: %s\n", result)
}
