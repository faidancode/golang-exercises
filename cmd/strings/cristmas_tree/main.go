package main

import (
	"fmt"
	"strings"
)

func main() {
	height := 10

	for i := 1; i <= height; i++ {
		indent := strings.Repeat(" ", height-i)
		stars := strings.Repeat("*", 2*i-1)
		fmt.Printf("%s%s\n", indent, stars)
	}

	trunkWidth := 3
	trunkHeight := 2
	truntIndent := strings.Repeat(" ", height-(trunkWidth/2)-1)

	for i := 0; i < trunkHeight; i++ {
		fmt.Printf("%s%s\n", truntIndent, strings.Repeat("#", trunkWidth))
	}
}
