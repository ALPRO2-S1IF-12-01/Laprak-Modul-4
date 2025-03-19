//103112400084
//Nufail Alauddin Tsaqif
package main

import "fmt"

func hitungSkor(soal int, totalwaktu *int, totalsoal *int) {
	if soal <= 300 {
		*totalsoal++
		*totalwaktu += soal
	}
}

func bandingkanPemain(nama1 string, totalsoal1, totalwaktu1 int, nama2 string, totalsoal2, totalwaktu2 int) (string, int, int) {
	if totalsoal1 > totalsoal2 || (totalsoal1 == totalsoal2 && totalwaktu1 < totalwaktu2) {
		return nama1, totalsoal1, totalwaktu1
	}
	return nama2, totalsoal2, totalwaktu2
}

func main() {
	var nama1, nama2 string
	var soal1, soal2 int
	var totalwaktu1, totalwaktu2, totalsoal1, totalsoal2 int
	var totalsoaltemp, totalwaktutemp int
	var namatemp string

	fmt.Scan(&nama1)
	if nama1 != "Selesai" {
		for i := 0; i < 8; i++ {
			fmt.Scan(&soal1)
			hitungSkor(soal1, &totalwaktu1, &totalsoal1)
		}

		for {
			fmt.Scan(&nama2)
			if nama2 == "Selesai" {
				break
			}

			totalwaktu2, totalsoal2 = 0, 0
			for i := 0; i < 8; i++ {
				fmt.Scan(&soal2)
				hitungSkor(soal2, &totalwaktu2, &totalsoal2)
			}

			namatemp, totalsoaltemp, totalwaktutemp = bandingkanPemain(nama1, totalsoal1, totalwaktu1, nama2, totalsoal2, totalwaktu2)
		}

		fmt.Println(namatemp, totalsoaltemp, totalwaktutemp)
	}
}
