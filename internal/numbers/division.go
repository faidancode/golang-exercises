package numbers

// Division mencari Faktor Persekutuan Terbesar (GCF/FPB) menggunakan algoritma Euclidean.
func Division(num1, num2 int) int {
	// Algoritma Euclidean Iteratif
	for num2 != 0 {
		// Simpan sisa bagi
		temp := num1 % num2

		// Geser nilai: num1 menjadi num2, num2 menjadi sisa baginya
		num1 = num2
		num2 = temp
	}

	// Ketika num2 mencapai 0, maka num1 adalah GCF-nya
	return num1
}
