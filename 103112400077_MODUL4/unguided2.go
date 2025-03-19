package main

import "fmt"

func hitungSkor(w1, w2, w3, w4, w5, w6, w7, w8 int, soalSelesai, totalWaktu *int) {
	*soalSelesai, *totalWaktu = 0, 0

	if w1 < 301 {
		*soalSelesai++
		*totalWaktu += w1
	}
	if w2 < 301 {
		*soalSelesai++
		*totalWaktu += w2
	}
	if w3 < 301 {
		*soalSelesai++
		*totalWaktu += w3
	}
	if w4 < 301 {
		*soalSelesai++
		*totalWaktu += w4
	}
	if w5 < 301 {
		*soalSelesai++
		*totalWaktu += w5
	}
	if w6 < 301 {
		*soalSelesai++
		*totalWaktu += w6
	}
	if w7 < 301 {
		*soalSelesai++
		*totalWaktu += w7
	}
	if w8 < 301 {
		*soalSelesai++
		*totalWaktu += w8
	}
}

func main() {
	var nama, pemenang string
	var w1, w2, w3, w4, w5, w6, w7, w8 int
	var soalSelesai, totalWaktu, maxSoal, minWaktu int
	maxSoal = -1
	minWaktu = 99999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		fmt.Scan(&w1, &w2, &w3, &w4, &w5, &w6, &w7, &w8)
		hitungSkor(w1, w2, w3, w4, w5, w6, w7, w8, &soalSelesai, &totalWaktu)

		if soalSelesai > maxSoal || (soalSelesai == maxSoal && totalWaktu < minWaktu) {
			maxSoal, minWaktu = soalSelesai, totalWaktu
			pemenang = nama
		}
	}

	fmt.Print(pemenang," ", maxSoal, minWaktu)
}