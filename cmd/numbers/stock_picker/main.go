package main

import (
	"fmt"
	"golang-exercises/internal/numbers"
)

func main() {
	input := []int{44, 30, 24, 32, 35, 30, 40, 38, 15}
	result := numbers.StockPicker(input)

	fmt.Printf("Input : %v\n", input)
	fmt.Printf("Result : %d\n", result) // Output: 16
}
