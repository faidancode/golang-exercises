package counter

import "strings"

func VowelCounter(str string) int {
	count := 0
	vowels := "aiueoAIUEO"

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
