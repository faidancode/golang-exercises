package reverser

func Reverse(str string) string {
	// Gunakan rune daripada byte: Di Go, string adalah sekumpulan byte. Jika input berisi karakter unik (seperti emoji atau bahasa Mandarin), membalikkan byte akan merusak karakter tersebut. Menggunakan []rune(s) adalah praktik terbaik.
	runes := []rune(str)

	// Mendeklarasikan dan menginisialisasi dua variabel dalam satu baris.
	// Tujuan: Teknik ini biasanya digunakan untuk algoritma Two Pointers. Dalam kasus string reverse, i akan bergerak maju (i++) dan j akan bergerak mundur (j--) hingga keduanya bertemu di tengah, sambil menukar nilai karakter di posisi tersebut.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
