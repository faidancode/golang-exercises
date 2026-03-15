package palindrome

import "unicode"

func IsPalindrome(str string) bool {
	runes := []rune(str)
	left := 0
	right := len(runes) - 1

	for left < right {
		if !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
			left++
			continue
		}

		if !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
			right--
			continue
		}

		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}

		left++
		right--

	}

	return true
}
