// PRATAMA BINTANG DANISWARA 103112400051
package main

import "fmt"

func hitungSkor() (int, int) {
	soal := 0
	skor := 0
	for i := 0; i < 8; i++ {
		var waktu int
		fmt.Scan(&waktu)
		if waktu < 301 {
			soal = soal + 1
			skor = skor + waktu
		}
	}
	return soal, skor
}

func main() {
	var nama string
	Soalterbanyak := 0
	Skorpalingsedikit := 99999
	var pemenang string

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		soal, skor := hitungSkor()
		if soal > Soalterbanyak || (soal == Soalterbanyak && skor < Skorpalingsedikit) {
			Soalterbanyak = soal
			Skorpalingsedikit = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, Soalterbanyak, Skorpalingsedikit)
}
