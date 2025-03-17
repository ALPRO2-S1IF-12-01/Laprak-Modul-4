// Felix Pedrosa V

package main

import (
	"fmt"
	"strings"
)

const waktuMaksimal = 301 // Waktu maksimal jika soal tidak selesai

// Fungsi untuk menghitung jumlah soal yang diselesaikan dan total waktu
func hitungNilai(waktuSoal [8]int, jumlahSoal *int, totalWaktu *int) {
	*jumlahSoal = 0
	*totalWaktu = 0
	for _, waktu := range waktuSoal {
		if waktu < waktuMaksimal { // Jika soal diselesaikan
			*jumlahSoal++
			*totalWaktu += waktu
		}
	}
}

func main() {
	var namaPemenang string
	var jumlahSoalPemenang, totalWaktuPemenang int
	jumlahSoalPemenang = 0
	totalWaktuPemenang = waktuMaksimal * 8 // Nilai awal skor pemenang maksimal

	for {
		var namaPeserta string
		var waktuPenyelesaian [8]int

		// Membaca input nama peserta
		fmt.Scan(&namaPeserta)
		namaPeserta = strings.TrimSpace(namaPeserta)

		// Jika input adalah 'Selesai', hentikan loop
		if namaPeserta == "Selesai" {
			break
		}

		// Membaca waktu penyelesaian 8 soal
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktuPenyelesaian[i])
		}

		// Hitung soal yang diselesaikan dan total waktu
		var jumlah, total int
		hitungNilai(waktuPenyelesaian, &jumlah, &total)

		// Tentukan pemenang berdasarkan jumlah soal dan total waktu
		if jumlah > jumlahSoalPemenang || (jumlah == jumlahSoalPemenang && total < totalWaktuPemenang) {
			namaPemenang = namaPeserta
			jumlahSoalPemenang = jumlah
			totalWaktuPemenang = total
		}
	}

	// Cetak pemenang
	fmt.Printf("%s %d %d\n",namaPemenang, jumlahSoalPemenang, totalWaktuPemenang)
}