package main

import (
	"fmt"
	"golang-exercises/internal/numbers"
)

func main() {
	input := 138
	result := numbers.TimeConvert(input)

	fmt.Printf("Input : %d\n", input)
	fmt.Printf("Result : %s\n", result)
}
