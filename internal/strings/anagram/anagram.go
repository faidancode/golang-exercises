package anagram

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Gunakan map untuk menghitung frekuensi karakter
	counts := make(map[rune]int)

	for _, r := range s1 {
		counts[r]++
	}

	for _, r := range s2 {
		counts[r]--
		// Jika hitungan menjadi negatif, artinya s2 punya karakter
		// yang tidak ada di s1 atau jumlahnya lebih banyak
		if counts[r] < 0 {
			return false
		}
	}

	return true

}
