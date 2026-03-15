package main

import (
	"fmt"
	"golang-exercises/internal/strings/anagram"
)

func main() {
	s1 := "listen"
	s2 := "silent"

	if anagram.IsAnagram(s1, s2) {
		fmt.Printf("'%s' dan '%s' ADALAH anagram\n", s1, s2)
	} else {
		fmt.Printf("'%s' dan '%s' BUKAN anagram\n", s1, s2)
	}

}
