package main

import (
	"fmt"
	"golang-exercises/internal/strings/palindrome"
)

func main() {
	words := []string{"Kodok", "A man, a plan, a canal: Panama", "Golang"}

	for _, w := range words {
		result := palindrome.IsPalindrome(w)
		fmt.Printf("Is %s palindrome? %v\n", w, result)
	}
}
