package main

import (
	"fmt"
)

// Fungsi untuk menghitung jumlah soal yang diselesaikan dan total waktu
func hitungSkor(waktu []int) (int, int) {
	batasWaktu := 301
	jumlahSoal, totalWaktu := 0, 0
	for _, t := range waktu {
		if t < batasWaktu {
			jumlahSoal++
			totalWaktu += t
		}
	}
	return jumlahSoal, totalWaktu
}

func main() {
	var pemenang string
	var maxSoal, minWaktu int = 0, 1000000

	for {
		var nama string
		var waktu [8]int
		_, err := fmt.Scan(&nama)
		if err != nil || nama == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			_, err = fmt.Scan(&waktu[i])
			if err != nil {
				break
			}
		}

		jumlahSoal, totalWaktu := hitungSkor(waktu[:])

		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minWaktu) {
			maxSoal = jumlahSoal
			minWaktu = totalWaktu
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}