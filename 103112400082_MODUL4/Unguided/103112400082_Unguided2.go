package main

import (
	"fmt"
)

func hitungSkor(waktu []int) (int, int) {
	totalSoal, totalWaktu := 0, 0
	for _, w := range waktu {
		if w < 301 {
			totalSoal++
			totalWaktu += w
		}
	}
	return totalSoal, totalWaktu
}

func main() {
	var pemenang string
	maxSoal, minWaktu := 0, 99999

	for {
		var nama string
		waktu := make([]int, 8)
		_, err := fmt.Scan(&nama)
		if err != nil || nama == "Selesai" {
			break
		}
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		soal, skor := hitungSkor(waktu)
		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang, maxSoal, minWaktu = nama, soal, skor
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
