package main

import "fmt"

func hitungSkor(jumlahSoal int, waktu *int, soal *int) {
	if jumlahSoal <= 300 {
		*soal += 1
		*waktu += jumlahSoal
	}
}

func main() {
	var nama1, nama2 string
	var soalInput int
	var totalWaktu, totalSoal int
	var maxSoal, minWaktu int
	var pemainWaktu, pemainSoal int
	fmt.Scan(&nama1)
	if nama1 == "Selesai" {
		return
	}
	nama2 = nama1
	for i := 0; i < 8; i++ {
		fmt.Scan(&soalInput)
		hitungSkor(soalInput, &totalWaktu, &totalSoal)
	}
	maxSoal = totalSoal
	minWaktu = totalWaktu
	for {
		fmt.Scan(&nama1)
		if nama1 == "Selesai" {
			break
		}
		pemainWaktu, pemainSoal = 0, 0
		for i := 0; i < 8; i++ {
			fmt.Scan(&soalInput)
			hitungSkor(soalInput, &pemainWaktu, &pemainSoal)
		}
		if pemainSoal > maxSoal || (pemainSoal == maxSoal && pemainWaktu < minWaktu) {
			nama2 = nama1
			maxSoal = pemainSoal
			minWaktu = pemainWaktu
		}
	}
	fmt.Println(nama2, maxSoal, minWaktu)
}
