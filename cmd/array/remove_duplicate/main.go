package main

import (
	"fmt"
	"golang-exercises/internal/array"
)

func main() {
	input := []int{1, 2, 3, 4, 4, 5, 5, 6, 7, 7, 7, 8}
	fmt.Printf("Original: %v\n", input)

	newLength := array.RemoveDuplicates(input)
	fmt.Printf("After Remove: %v, New Length: %d\n", input[:newLength], newLength)
}
