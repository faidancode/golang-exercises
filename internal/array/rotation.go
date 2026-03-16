package array

import (
	"slices"
	"strconv"
	"strings"
)

// Circular Rotation (rotasi melingkar) pada array. Aturannya sederhana: ambil elemen pertama sebagai indeks penentu ($N$), lalu putar array sehingga elemen pada indeks ke-$N$ menjadi elemen pertama di hasil akhir.
func ArrayRotation(arr []int) string {
	if len(arr) == 0 {
		return ""
	}

	n := arr[0] % len(arr)

	rotated := slices.Concat(arr[n:], arr[:n])

	var result strings.Builder
	for _, val := range rotated {
		result.WriteString((strconv.Itoa(val)))
	}

	return result.String()
}
