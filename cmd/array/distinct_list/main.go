package main

import (
	"fmt"
	"golang-exercises/internal/array"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 5, 6, 6, 7, 8}
	result := array.DistinctList(input)

	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Result: %d\n", result)
}
