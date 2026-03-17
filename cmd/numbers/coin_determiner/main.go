package main

import (
	"fmt"
	"golang-exercises/internal/numbers"
)

func main() {
	input := 17
	result := numbers.CoinDeterminer(input)

	fmt.Printf("Input : %v\n", input)
	fmt.Printf("Result : %v", result)
}
