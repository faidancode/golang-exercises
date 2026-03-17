package numbers

// CoinDeterminer mencari jumlah koin minimum untuk mencapai target num.
func CoinDeterminer(num int) int {
	// Coin Change Problem. Karena kita ingin mencari jumlah koin paling sedikit (minimum coins), pendekatan terbaik adalah menggunakan Dynamic Programming (DP).

	coins := []int{1, 5, 7, 9, 11}

	// Buat tabel DP untuk menyimpan jumlah koin minimum untuk setiap nilai 0 sampai num
	// Kita isi dengan nilai yang sangat besar (infinity) sebagai inisialisasi
	dp := make([]int, num+1)
	for i := range dp {
		dp[i] = num + 1 // num + 1 pasti lebih besar dari kemungkinan kombinasi koin terkecil
	}

	// Base case: untuk mencapai nilai 0, kita butuh 0 koin
	dp[0] = 0

	// Iterasi untuk setiap nilai dari 1 sampai num
	for i := 1; i <= num; i++ {
		// Cek setiap jenis koin yang tersedia
		for _, coin := range coins {
			if i >= coin {
				// Rumus DP: dp[i] = min(dp[i], dp[i - coin] + 1)
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	// Pendekatan ini bekerja dengan cara membangun solusi untuk nilai kecil terlebih dahulu, lalu menggunakannya untuk menghitung nilai yang lebih besar:

	// Tabel dp: dp[i] menyimpan jumlah koin paling sedikit untuk menghasilkan angka i.
	// Transisi: Untuk menghitung dp[16], kita mengecek semua kemungkinan koin terakhir yang digunakan:
	// Jika koin terakhir 11, maka kita butuh 1 + dp[5] koin.
	// Jika koin terakhir 9, maka kita butuh 1 + dp[7] koin.
	// Jika koin terakhir 7, maka kita butuh 1 + dp[9] koin.
	// ... dan seterusnya.
	// Hasil: Nilai terkecil dari pilihan-pilihan di atas akan disimpan di dp[16].

	return dp[num]
}
