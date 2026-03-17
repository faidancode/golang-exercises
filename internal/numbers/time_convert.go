package numbers

import (
	"fmt"
)

// TimeConvert mengubah total menit menjadi format jam:menit.
func TimeConvert(num int) string {
	// Menghitung jam dengan pembagian bulat
	hours := num / 60

	// Menghitung sisa menit dengan operator modulo
	minutes := num % 60

	// Menggabungkan keduanya dalam format string "jam:menit"
	return fmt.Sprintf("%d:%d", hours, minutes)
}
