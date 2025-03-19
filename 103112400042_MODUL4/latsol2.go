// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func hitungSkor(waktu int, totalWaktu, totalSoal *int) bool {
	if waktu <= 300 {
		*totalSoal++
		*totalWaktu += waktu
		return true
	}
	return false
}
func hitungSkorPeserta(namaPeserta string, waktu []int) (int, int) {
	var totalWaktu, totalSoal int
	for _, w := range waktu {
		if hitungSkor(w, &totalWaktu, &totalSoal) {
			// tidak ada error
		} else {
			fmt.Printf("Waktu %d peserta %s lebih dari 300\n", w, namaPeserta)
		}
	}
	return totalSoal, totalWaktu
}

func main() {
	var namaPemenang string
	var maxSoal, maxWaktu int

	var namaPeserta1 string
	fmt.Print("Masukkan nama peserta 1: ")
	fmt.Scanln(&namaPeserta1)

	var waktuPeserta1 [8]int
	for i := 0; i < 8; i++ {
		fmt.Printf("Masukkan waktu soal %d peserta 1: ", i+1)
		fmt.Scanln(&waktuPeserta1[i])
	}

	totalSoal1, totalWaktu1 := hitungSkorPeserta(namaPeserta1, waktuPeserta1[:])
	var namaPeserta2 string
	fmt.Print("Masukkan nama peserta 2: ")
	fmt.Scanln(&namaPeserta2)

	var waktuPeserta2 [8]int
	for i := 0; i < 8; i++ {
		fmt.Printf("Masukkan waktu soal %d peserta 2: ", i+1)
		fmt.Scanln(&waktuPeserta2[i])
	}

	totalSoal2, totalWaktu2 := hitungSkorPeserta(namaPeserta2, waktuPeserta2[:])

	if totalSoal1 > totalSoal2 {
		namaPemenang = namaPeserta1
		maxSoal = totalSoal1
		maxWaktu = totalWaktu1
	} else if totalSoal2 > totalSoal1 {
		namaPemenang = namaPeserta2
		maxSoal = totalSoal2
		maxWaktu = totalWaktu2
	} else {
		if totalWaktu1 < totalWaktu2 {
			namaPemenang = namaPeserta1
			maxSoal = totalSoal1
			maxWaktu = totalWaktu1
		} else {
			namaPemenang = namaPeserta2
			maxSoal = totalSoal2
			maxWaktu = totalWaktu2
		}
	}

	fmt.Printf("Pemenang: %s, Skor: %d, Waktu: %d\n", namaPemenang, maxSoal, maxWaktu)
}
