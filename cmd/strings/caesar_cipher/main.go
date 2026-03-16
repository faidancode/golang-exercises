package main

import (
	"fmt"
	caesarcipher "golang-exercises/internal/strings/caesar_cipher"
)

func main() {
	input := "ABCDPQRS"
	num := 3
	fmt.Println(caesarcipher.CaesarCipher(input, num))
}
