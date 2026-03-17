package numbers

import (
	"sort"
	"strconv"
)

// PermutationStep mencari permutasi angka berikutnya yang lebih besar.
func PermutationStep(num int) int {
	// 1. Ubah angka menjadi slice of digits
	s := strconv.Itoa(num)
	digits := []rune(s)
	n := len(digits)

	// 2. Cari 'pivot' (digit pertama dari kanan yang lebih kecil dari kanannya)
	pivot := -1
	for i := n - 2; i >= 0; i-- {
		if digits[i] < digits[i+1] {
			pivot = i
			break
		}
	}

	// Jika tidak ada pivot, berarti angka sudah maksimal
	if pivot == -1 {
		return -1
	}

	// 3. Cari digit terkecil di kanan pivot yang lebih besar dari pivot
	successor := -1
	for i := n - 1; i > pivot; i-- {
		if digits[i] > digits[pivot] {
			successor = i
			break
		}
	}

	// 4. Tukar pivot dengan successor
	digits[pivot], digits[successor] = digits[successor], digits[pivot]

	// 5. Urutkan angka di sebelah kanan pivot agar menjadi angka terkecil
	// digits[pivot+1:] (Slicing)
	rightSide := digits[pivot+1:]
	sort.Slice(rightSide, func(i, j int) bool {
		return rightSide[i] < rightSide[j]
	})

	// Gabungkan kembali dan ubah ke integer
	resultStr := string(digits)
	result, _ := strconv.Atoi(resultStr)

	return result
}
