package main

import (
	"fmt"
	"strings"
)

func hitungSkor() (int, int) {
	var jumlahSoal, totalWaktu, waktu int
	jumlahSoal, totalWaktu = 0, 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu < 301 {
			jumlahSoal++
			totalWaktu += waktu
		}
	}

	return jumlahSoal, totalWaktu
}

func main() {
	var pemenang string
	var soalTerbanyak, skorTerendah int
	soalTerbanyak, skorTerendah = -1, 99999999

	fmt.Println("Masukkan data peserta (ketik 'Selesai' untuk berhenti):")

	for {
		var nama string
		fmt.Scan(&nama)

		if strings.ToLower(nama) == "selesai" {
			break
		}

		jumlahSoal, totalWaktu := hitungSkor()

		if jumlahSoal > soalTerbanyak || (jumlahSoal == soalTerbanyak && totalWaktu < skorTerendah) {
			pemenang = nama
			soalTerbanyak = jumlahSoal
			skorTerendah = totalWaktu
		}
	}

	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, soalTerbanyak, skorTerendah)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
