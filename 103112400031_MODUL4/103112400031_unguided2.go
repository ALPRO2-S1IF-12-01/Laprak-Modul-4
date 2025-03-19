// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func hitungSkor(waktuPengerjaan []int) (int, int) {
	soalSelesai, totalDurasi := 0, 0

	for _, durasi := range waktuPengerjaan {
		if durasi <= 300 {
			soalSelesai++
			totalDurasi += durasi
		}
	}
	return soalSelesai, totalDurasi
} 

func main() {
	var namaPemenang string
	var soalTerbanyak, waktuTercepat int = 0, 1000000

	for {
		var namaPeserta string
		var waktuPengerjaan [8]int

		_, err := fmt.Scan(&namaPeserta)
		if err != nil || namaPeserta == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			_, err = fmt.Scan(&waktuPengerjaan[i])
			if err != nil {
				break
			}
		}

		soalSelesai, totalDurasi := hitungSkor(waktuPengerjaan[:])

		if soalSelesai > soalTerbanyak || (soalSelesai == soalTerbanyak && totalDurasi < waktuTercepat) {
			namaPemenang = namaPeserta
			soalTerbanyak = soalSelesai
			waktuTercepat = totalDurasi
		}
	}
	fmt.Println(namaPemenang, soalTerbanyak, waktuTercepat)
}
