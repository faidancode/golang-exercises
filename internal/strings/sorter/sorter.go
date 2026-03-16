package sorter

import (
	"slices"
)

func Sorter(str string) string {
	runes := []rune(str)

	slices.Sort(runes)

	return string(runes)
}
