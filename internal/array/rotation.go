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
	// Menentukan Titik Potong: Jika array adalah [2, 3, 4, 1, 6, 10], maka $N = 2$.
	// Artinya, elemen pada indeks ke-2 (angka 4) akan menjadi pemimpin barisan yang baru.
	n := arr[0] % len(arr)

	// Slicing di Go:
	// arr[n:] mengambil elemen dari indeks $N$ hingga akhir: [4, 1, 6, 10].
	// arr[:n] mengambil elemen dari awal hingga sebelum indeks $N$: [2, 3].
	rotated := slices.Concat(arr[n:], arr[:n])

	// Konversi String: Kita menggunakan strings.Builder untuk menggabungkan angka-angka menjadi satu string. Ini jauh lebih cepat daripada menggunakan operator + berulang kali di dalam loop.
	var result strings.Builder
	for _, val := range rotated {
		result.WriteString((strconv.Itoa(val)))
	}

	return result.String()
}
