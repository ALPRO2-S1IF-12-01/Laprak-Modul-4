// Raihan Adi Arba
// 103112400071
package main

import "fmt"

func main() {
	var (
		winnerNama                                                     string
		winnerSkor                                                     = 0
		winnerSoal                                                     = 0
		nama                                                           string
		waktu1, waktu2, waktu3, waktu4, waktu5, waktu6, waktu7, waktu8 int
		soal, skor                                                     int
	)

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		fmt.Scan(&waktu1, &waktu2, &waktu3, &waktu4, &waktu5, &waktu6, &waktu7, &waktu8)

		hitungSkor(waktu1, waktu2, waktu3, waktu4, waktu5, waktu6, waktu7, waktu8, &soal, &skor)

		if soal > winnerSoal || (soal == winnerSoal && skor < winnerSkor) {
			winnerNama = nama
			winnerSkor = skor
			winnerSoal = soal
		}
	}

	fmt.Println(winnerNama, winnerSoal, winnerSkor)
}
func hitungSkor(w1, w2, w3, w4, w5, w6, w7, w8 int, soal, skor *int) {
	*soal = 0
	*skor = 0

	if w1 <= 300 {
		*soal += 1
		*skor += w1
	}
	if w2 <= 300 {
		*soal += 1
		*skor += w2
	}
	if w3 <= 300 {
		*soal += 1
		*skor += w3
	}
	if w4 <= 300 {
		*soal += 1
		*skor += w4
	}
	if w5 <= 300 {
		*soal += 1
		*skor += w5
	}
	if w6 <= 300 {
		*soal += 1
		*skor += w6
	}
	if w7 <= 300 {
		*soal += 1
		*skor += w7
	}
	if w8 <= 300 {
		*soal += 1
		*skor += w8
	}
}
