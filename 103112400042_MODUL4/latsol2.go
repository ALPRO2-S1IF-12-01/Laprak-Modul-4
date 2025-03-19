package main

import "fmt"

// Fungsi untuk menghitung skor
func hitungSkor(waktu int, totalWaktu, totalSoal *int) bool {
	if waktu <= 300 {
		*totalSoal++
		*totalWaktu += waktu
		return true
	}
	return false
}

// Fungsi untuk menghitung skor peserta
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

	// Input nama peserta 1
	var namaPeserta1 string
	fmt.Print("Masukkan nama peserta 1: ")
	fmt.Scanln(&namaPeserta1)

	// Input waktu peserta 1
	var waktuPeserta1 [8]int
	for i := 0; i < 8; i++ {
		fmt.Printf("Masukkan waktu soal %d peserta 1: ", i+1)
		fmt.Scanln(&waktuPeserta1[i])
	}

	// Hitung skor peserta 1
	totalSoal1, totalWaktu1 := hitungSkorPeserta(namaPeserta1, waktuPeserta1[:])

	// Input nama peserta 2
	var namaPeserta2 string
	fmt.Print("Masukkan nama peserta 2: ")
	fmt.Scanln(&namaPeserta2)

	// Input waktu peserta 2
	var waktuPeserta2 [8]int
	for i := 0; i < 8; i++ {
		fmt.Printf("Masukkan waktu soal %d peserta 2: ", i+1)
		fmt.Scanln(&waktuPeserta2[i])
	}

	// Hitung skor peserta 2
	totalSoal2, totalWaktu2 := hitungSkorPeserta(namaPeserta2, waktuPeserta2[:])

	// Tentukan pemenang
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
