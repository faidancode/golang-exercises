package numbers

// FibonacciChecker mengecek apakah num ada di dalam deret Fibonacci.
func FibonacciChecker(num int) string {
	// Base case: 0 dan 1 adalah bagian dari Fibonacci
	if num == 0 || num == 1 {
		return "yes"
	}

	a, b := 0, 1

	// Terus hasilkan angka Fibonacci sampai mencapai atau melewati num
	for b < num {
		// Simpan nilai b sementara, lalu update a dan b
		// a, b = b, a+b adalah cara singkat di Go
		temp := b
		b = a + b
		a = temp
	}

	// Jika b sama dengan num, berarti num adalah angka Fibonacci
	if b == num {
		return "yes"
	}

	// Deret Fibonacci bekerja dengan menjumlahkan dua angka terakhir:0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
	// Iterasi: Kita mulai dengan dua angka pertama, 0 dan 1.
	// Looping: Di dalam for, kita terus menghitung angka berikutnya dengan menjumlahkan dua angka sebelumnya.
	// Pengecekan:
	// - Jika saat proses penjumlahan kita menemukan angka yang sama dengan num, kita kembalikan "yes".
	// - Jika angka yang kita hasilkan sudah melebihi num tanpa pernah menyamainya, berarti num bukan bagian dari Fibonacci, maka kita kembalikan "no".

	return "no"
}
