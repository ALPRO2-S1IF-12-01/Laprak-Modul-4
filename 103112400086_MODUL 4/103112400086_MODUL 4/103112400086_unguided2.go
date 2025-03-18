package main

import "fmt"

func skor(durasi [8]int) (int, int) {
	jumlahSoal := 0
	totalWaktu := 0

	for _, waktu := range durasi {
		if waktu < 301 {
			jumlahSoal++
			totalWaktu = totalWaktu + waktu
		}
	}
	return jumlahSoal, totalWaktu
}

func main() {
	var winner string

	maxSoal := -1      // -1 biar nilai prtama yg dibaca iterasi lbh gede
	minWaktu := 999999 // dikasih angka bsar diawal biar nilai prtama yg dibaca lbh kecil

	for {
		var nama string

		fmt.Scan(&nama)

		if nama == "Selesai" {
			break

		}

		var durasi [8]int

		for i := range durasi {
			fmt.Scan(&durasi[i])
		}

		jumlahSoal, totalWaktu := skor(durasi)

		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minWaktu) {
			winner, maxSoal, minWaktu = nama, jumlahSoal, totalWaktu
		}
	}

	fmt.Println(winner, maxSoal, minWaktu)
}
