package main

import (
	"fmt"
)

func hitungSkor(waktu [8]int) (int, int) {
	jumlahSoal, totalWaktu := 0, 0
	for _, w := range waktu {
		if w <= 300 {
			jumlahSoal++
			totalWaktu += w
		}
	}
	return jumlahSoal, totalWaktu
}

func main() {
	var nama, pemenang string
	var waktu [8]int
	maxSoal, minWaktu := 0, 999999

	for {
		fmt.Print("Masukkan nama peserta (atau 'Selesai' untuk berhenti): ")
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		soal, total := hitungSkor(waktu)

		if soal > maxSoal || (soal == maxSoal && total < minWaktu) {
			pemenang, maxSoal, minWaktu = nama, soal, total
		}
	}

	if pemenang != "" {
		fmt.Println(pemenang, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta yang valid.")
	}
}
