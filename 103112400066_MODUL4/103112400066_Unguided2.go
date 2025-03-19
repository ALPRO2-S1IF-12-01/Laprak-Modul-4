// DWI OKTA SURYANINGRUM
package main

import (
	"fmt"
)

// Fungsi hitungSkor digunakan untuk menghitung berapa banyak soal yang berhasil diselesaikan dan total waktu yang dibutuhkan oleh seorang peserta.
// Parameter:
// - waktuPengerjaan: Daftar waktu yang dibutuhkan untuk menyelesaikan setiap soal.
// - totalSoal: Variabel untuk menyimpan jumlah soal yang berhasil diselesaikan.
// - totalSkor: Variabel untuk menyimpan total waktu yang dibutuhkan.
func hitungSkor(waktuPengerjaan []int, totalSoal *int, totalSkor *int) {
	// Mulai dengan mengatur totalSoal dan totalSkor ke 0
	*totalSoal = 0
	*totalSkor = 0

	// Loop melalui setiap waktu pengerjaan
	for _, waktu := range waktuPengerjaan {
		// Jika waktu pengerjaan kurang dari atau sama dengan 300 menit (5 jam), artinya soal tersebut berhasil diselesaikan.
		if waktu <= 300 {
			*totalSoal++ // Tambahkan 1 ke totalSoal
			*totalSkor += waktu // Tambahkan waktu ke totalSkor
		}
	}
}

func main() {
	// Variabel untuk menyimpan nama pemenang, jumlah soal terbanyak, dan total waktu terkecil
	var namaPemenang string
	var maxSoal, minSkor int

	// Loop utama untuk meminta input dari user
	for {
		var nama string
		// Meminta user memasukkan nama peserta
		fmt.Print("Masukkan nama peserta (atau ketik 'Selesai' untuk mengakhiri): ")
		fmt.Scan(&nama)

		// Jika user mengetik "Selesai", maka loop akan dihentikan
		if nama == "Selesai" {
			break
		}

		// Array untuk menyimpan waktu pengerjaan 8 soal
		var waktuPengerjaan [8]int
		// Meminta user untuk memasukkan waktu pengerjaan untuk 8 soal
		fmt.Print("Masukkan waktu pengerjaan untuk 8 soal (dalam menit): ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktuPengerjaan[i]) // Membaca waktu pengerjaan untuk setiap soal
		}

		// Variabel untuk menyimpan total soal dan total skor peserta
		var totalSoal, totalSkor int
		// Panggil fungsi hitungSkor untuk menghitung total soal dan total skor
		hitungSkor(waktuPengerjaan[:], &totalSoal, &totalSkor)

		// Bandingkan skor peserta dengan skor tertinggi saat ini
		// Jika peserta ini menyelesaikan lebih banyak soal, atau
		// jika jumlah soalnya sama tetapi waktu pengerjaannya lebih cepat,
		// maka peserta ini menjadi pemenang sementara.
		if totalSoal > maxSoal || (totalSoal == maxSoal && totalSkor < minSkor) {
			namaPemenang = nama // Simpan nama pemenang
			maxSoal = totalSoal // Simpan jumlah soal terbanyak
			minSkor = totalSkor // Simpan total waktu terkecil
		}
	}

	// Setelah loop selesai, tampilkan hasil pemenang
	if namaPemenang != "" {
		fmt.Printf("Pemenang: %s, Jumlah soal: %d, Nilai: %d\n", namaPemenang, maxSoal, minSkor)
	} else {
		// Jika tidak ada peserta yang masuk, tampilkan pesan ini
		fmt.Println("Tidak ada peserta yang masuk.")
	}
}