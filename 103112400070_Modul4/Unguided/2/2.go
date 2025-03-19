// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func main() {
	var maxDurasi = 301
	var pemenang string
	var maxSoal, minWaktu int
	maxSoal = 0
	minWaktu = maxDurasi * 8

	for {
		var peserta string
		var durasi [8]int
		fmt.Scan(&peserta)

		if peserta == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&durasi[i])
		}

		var soalTerjawab, totalDurasi int
		hitungNilai(durasi, &soalTerjawab, &totalDurasi, maxDurasi)

		if soalTerjawab > maxSoal || (soalTerjawab == maxSoal && totalDurasi < minWaktu) {
			pemenang = peserta
			maxSoal = soalTerjawab
			minWaktu = totalDurasi
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}

func hitungNilai(durasi [8]int, soal *int, total *int, maxDurasi int) {
	*soal = 0
	*total = 0
	for _, waktu := range durasi {
		if waktu < maxDurasi {
			*soal++
			*total += waktu
		}
	}
}
