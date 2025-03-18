package main

import (
	"fmt"
	"strings"
)

func hitungSkor(waktu [8]int) (int, int) {
	jumlahSoal, totalWaktu := 0, 0
	for _, t := range waktu {
		if t <= 300 {
			jumlahSoal++
			totalWaktu += t
		}
	}
	return jumlahSoal, totalWaktu
}

func main() {
	var nama, pemenang string
	var waktu [8]int
	maxSoal, minWaktu := 0, 99999

	fmt.Println("Masukkan peserta dan waktu (akhiri dengan 'Selesai'): ")

	for {
		fmt.Scan(&nama)
		if strings.ToLower(nama) == "selesai" {
			break
		}
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		jumlahSoal, totalWaktu := hitungSkor(waktu)

		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minWaktu) {
			pemenang, maxSoal, minWaktu = nama, jumlahSoal, totalWaktu
		}
	}

	fmt.Printf("\nPemenang: %s\nJumlah Soal: %d\nTotal Waktu: %d menit\n", pemenang, maxSoal, minWaktu)
}
