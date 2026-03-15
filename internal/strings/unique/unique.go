package unique

func FirstUniqChar(str string) int {
	counts := make(map[rune]int)

	for _, r := range str {
		counts[r]++
	}

	for i, r := range str {
		if counts[r] == 1 {
			return i
		}
	}

	return -1
}
