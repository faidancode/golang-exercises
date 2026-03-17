package numbers

// StockPicker menghitung keuntungan maksimal dari deretan harga saham.
func StockPicker(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	// Inisialisasi harga beli terendah dengan hari pertama
	minPrice := arr[0]
	maxProfit := -1

	for i := 1; i < len(arr); i++ {
		currentPrice := arr[i]

		// Jika harga saat ini bisa memberikan profit yang lebih tinggi
		if currentPrice > minPrice {
			profit := currentPrice - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}

		// Update harga beli terendah jika kita menemukan harga yang lebih murah
		if currentPrice < minPrice {
			minPrice = currentPrice
		}
	}

	// Satu Kali Iterasi (O(n)): Kita tidak perlu membandingkan setiap pasang angka dengan dua loop bersarang. Itu akan lambat (O(n^2)).
	// minPrice: Kita terus memantau "kapan waktu termurah untuk membeli" yang pernah ada sebelum hari ini.
	// Kalkulasi Profit: Di setiap langkah, kita membayangkan: "Jika saya beli di harga paling murah sebelumnya (minPrice) dan jual hari ini (currentPrice), berapa untungnya?"
	// Update Profit: Jika untung hari ini lebih besar dari rekor maxProfit sebelumnya, kita perbarui rekornya.
	// Kasus Rugi: Jika sepanjang iterasi harga selalu turun (seperti [10, 9, 8]), maka maxProfit akan tetap -1

	return maxProfit
}
