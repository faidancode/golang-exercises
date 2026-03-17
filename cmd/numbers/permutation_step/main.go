package main

import (
	"fmt"
	"golang-exercises/internal/numbers"
)

func main() {
	input := 41352
	result := numbers.PermutationStep(input)

	// Cari Pivot: Dari kanan, 2 < 5 (ya), 5 < 3 (tidak), 3 < 1 (tidak), 3 > 1 (ya!). Jadi 3 adalah pivot kita (indeks ke-2).
	// Cari Successor: Di sebelah kanan 3, angka mana yang lebih besar dari 3 tapi paling kecil? Ada 5.
	// Tukar: Tukar 3 dengan 5. Angka menjadi 41532.
	// Urutkan Kanan: Angka di sebelah kanan posisi pivot tadi adalah 3, 2. Kita urutkan menjadi 2, 3.
	// Hasil Akhir: 41523.

	fmt.Printf("Input %v\n", input)
	fmt.Printf("Result %d\n", result)
}
