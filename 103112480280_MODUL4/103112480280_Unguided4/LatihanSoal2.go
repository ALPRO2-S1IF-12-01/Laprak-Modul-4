// Nama : Anggun Wahyu W. (103112480280)
package main

import "fmt"

// Prosedur untuk menghitung skor peserta
func hitungSkor(soal [8]int, totalSoal *int, totalWaktu *int) {
	*totalSoal = 0
	*totalWaktu = 0
	for _, waktu := range soal {
		if waktu <= 301 { // Jika waktu kurang dari atau sama dengan 301 menit
			*totalSoal++
			*totalWaktu += waktu
		}
	}
}

func main() {
	var nama1, nama2 string
	var soal1, soal2 [8]int
	var totalSoal1, totalWaktu1, totalSoal2, totalWaktu2 int

	// Input data peserta pertama
	fmt.Print("Masukkan nama peserta pertama: ")
	fmt.Scan(&nama1)
	fmt.Print("Masukkan waktu pengerjaan soal (8 angka):")
	for i := 0; i < 8; i++ {
		fmt.Scan(&soal1[i])
	}

	// Hitung skor peserta pertama
	hitungSkor(soal1, &totalSoal1, &totalWaktu1)

	// Input data peserta kedua
	fmt.Print("Masukkan nama peserta kedua: ")
	fmt.Scan(&nama2)
	fmt.Print("Masukkan waktu pengerjaan soal (8 angka):")
	for i := 0; i < 8; i++ {
		fmt.Scan(&soal2[i])
	}

	// Hitung skor peserta kedua
	hitungSkor(soal2, &totalSoal2, &totalWaktu2)

	// Menentukan pemenang
	if totalSoal1 > totalSoal2 || (totalSoal1 == totalSoal2 && totalWaktu1 < totalWaktu2) {
		fmt.Printf("%s %d %d\n", nama1, totalSoal1, totalWaktu1)
	} else {
		fmt.Printf("%s %d %d\n", nama2, totalSoal2, totalWaktu2)
	}
}
