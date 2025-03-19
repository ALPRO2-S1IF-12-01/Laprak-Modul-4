// KEISHIN NAUFA ALFARIDZHI
// 103112400061
package main

import "fmt"

func main() {
	var (
		winnerNama, nama string
		winnerSkor int = 0
		winnerSoal int = 0
		waktu [8]int
		soal, skor int
	)

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		hitungSkor(waktu, &soal, &skor)

		if soal > winnerSoal || (soal == winnerSoal && skor < winnerSkor) {
			winnerNama = nama
			winnerSkor = skor
			winnerSoal = soal
		}
	}

	fmt.Println(winnerNama, winnerSoal, winnerSkor)
}

func hitungSkor(waktu [8]int, soal, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if waktu[i] <= 300 {
			*soal += 1
			*skor += waktu[i]
		}
	}
}