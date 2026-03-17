package array

// DistinctList menghitung jumlah total entri duplikat dalam array.
func DistinctList(arr []int) int {
	// Map untuk menyimpan angka yang sudah pernah muncul
	seen := make(map[int]bool)
	duplicates := 0

	for _, num := range arr {
		// Jika angka sudah ada di dalam map 'seen'
		if seen[num] {
			// Berarti ini adalah kemunculan ke-2, ke-3, dst.
			duplicates++
		} else {
			// Jika belum ada, tandai angka ini sebagai sudah terlihat
			seen[num] = true
		}
	}

	return duplicates
}
