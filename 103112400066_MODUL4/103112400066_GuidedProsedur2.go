//DWI OKTA SURYANINGRUM
package main

import "fmt"

// Fungsi hitungRataRata menghitung rata-rata nilai mahasiswa dan menentukan status kelulusan.
// Parameter:
// - nama: Nama mahasiswa.
// - nilai1, nilai2, nilai3: Nilai-nilai mahasiswa.
func hitungRataRata(nama string, nilai1, nilai2, nilai3 float64) {
	// Hitung rata-rata nilai
	rataRata := (nilai1 + nilai2 + nilai3) / 3

	// Tentukan status kelulusan
	status := "Tidak Lulus" // Default status adalah "Tidak Lulus"
	if rataRata >= 60 {     // Jika rata-rata >= 60, status diubah menjadi "Lulus"
		status = "Lulus"
	}

	// Cetak hasil nilai akademik
	fmt.Println("\n=== Nilai Akademik ===")
	fmt.Println("Nama Mahasiswa : ", nama)
	fmt.Printf("Nilai 1		: %.2f\n", nilai1)
	fmt.Printf("Nilai 2		: %.2f\n", nilai2)
	fmt.Printf("Nilai 3		: %.2f\n", nilai3)
	fmt.Printf("Rata-rata		: %.2f\n", rataRata)
	fmt.Println("Status: ", status)
}

func main() {
	// Variabel untuk menyimpan input dari pengguna
	var nama string
	var nilai1, nilai2, nilai3 float64

	// Minta pengguna memasukkan nama mahasiswa
	fmt.Print("Masukkan Nama Mahasiswa : ")
	fmt.Scanln(&nama)

	// Minta pengguna memasukkan nilai 1
	fmt.Print("Masukkan Nilai 1 : ")
	fmt.Scanln(&nilai1)

	// Minta pengguna memasukkan nilai 2
	fmt.Print("Masukkan Nilai 2 : ")
	fmt.Scanln(&nilai2)

	// Minta pengguna memasukkan nilai 3
	fmt.Print("Masukkan Nilai 3 : ")
	fmt.Scanln(&nilai3)

	// Panggil fungsi hitungRataRata dengan data yang dimasukkan pengguna
	hitungRataRata(nama, nilai1, nilai2, nilai3)
}