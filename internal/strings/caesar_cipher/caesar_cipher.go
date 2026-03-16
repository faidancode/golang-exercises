package caesarcipher

import (
	"unicode"
)

// CaesarCipher menggeser setiap huruf sebanyak num posisi.
func CaesarCipher(str string, num int) string {
	runes := []rune(str)

	// Pastikan num tetap dalam rentang 0-25 dengan operator modulo
	// Modulo % 26: Ini sangat penting. Jika num adalah 27, itu sama saja dengan menggeser 1 kali ($27 \pmod{26} = 1$). Ini mencegah overflow dan membuat fungsi tetap bekerja meskipun input num sangat besar.
	shift := num % 26

	for i, char := range runes {
		if unicode.IsLetter(char) {
			// Tentukan titik awal ('a' untuk huruf kecil, 'A' untuk huruf besar)
			base := 'a'
			if unicode.IsUpper(char) {
				base = 'A'
			}

			// Rumus geser: (posisi_saat_ini + geseran - titik_awal) % 26 + titik_awal
			// Kita konversi ke int32 agar operasi matematikanya aman
			shiftedChar := (int(char)+shift-int(base))%26 + int(base)
			runes[i] = rune(shiftedChar)
		}
		// Karakter selain huruf (spasi, tanda baca) akan dibiarkan tetap.
	}

	return string(runes)
}
