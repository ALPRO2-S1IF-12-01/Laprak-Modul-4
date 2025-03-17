// Dimas Ramadhani
// 103112400065

package main

import "fmt"

func hitungSkor(soal int, totalwaktu *int, totalsoal *int) {
	if soal <= 300 {
		*totalsoal++
		*totalwaktu += soal
	}

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
			for i := 0; i < 8; i++ {
				fmt.Scan(&soal2)
				hitungSkor(soal2, &totalwaktu2, &totalsoal2)
			}
			if totalsoal1 < totalsoal2 {
				totalsoaltemp = totalsoal2
				totalwaktutemp = totalwaktu2
				namatemp = nama2
			} else if totalsoal1 > totalsoal2 {
				totalsoaltemp = totalsoal1
				totalwaktutemp = totalwaktu1
				namatemp = nama1
			} else if totalsoal1 == totalsoal2 {
				if totalwaktu1 > totalwaktu2 {
					totalsoaltemp = totalsoal1
					totalwaktutemp = totalwaktu1
					namatemp = nama1
				} else if totalsoal1 < totalsoal2 {
					totalsoaltemp = totalsoal2
					totalwaktutemp = totalwaktu2
					namatemp = nama2
				}
			}
		}
		fmt.Println(namatemp, totalsoaltemp, totalwaktutemp)
	}
	
}
