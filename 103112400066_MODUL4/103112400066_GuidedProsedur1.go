// DWI OKTA SURYANINGRUM
package main

import "fmt"

// Fungsi hitungGaji menghitung total gaji karyawan berdasarkan gaji pokok dan jam lembur.
// Parameter:
// - nama: Nama karyawan.
// - gajipokok: Gaji pokok karyawan.
// - jamLembur: Jumlah jam lembur yang dilakukan karyawan.
func hitungGaji(nama string, gajipokok float64, jamLembur int) {
	// Hitung bonus lembur: Rp50.000 per jam
	bonusLembur := float64(jamLembur) * 50000

	// Hitung total gaji: gaji pokok + bonus lembur
	totalGaji := gajipokok + bonusLembur

	// Cetak slip gaji
	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("")
	fmt.Println("Nama Karyawan :", nama)
	fmt.Printf("Gaji Pokok : Rp %2.f\n", gajipokok)
	fmt.Printf("Bonus Lembur : Rp %2.f(%d jam x Rp50.000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji : Rp %2.f\n", totalGaji)
}

func main() {
	// Variabel untuk menyimpan input dari pengguna
	var nama string
	var gajipokok float64
	var jamLembur int

	// Minta pengguna memasukkan nama karyawan
	fmt.Print("Masukkan Nama Karyawan : ")
	fmt.Scan(&nama)

	// Minta pengguna memasukkan gaji pokok
	fmt.Print("Masukkan Gaji Pokok : ")
	fmt.Scan(&gajipokok)

	// Minta pengguna memasukkan jumlah jam lembur
	fmt.Print("Masukkan Jumlah Jam Lembur : ")
	fmt.Scan(&jamLembur)

	// Panggil fungsi hitungGaji dengan data yang dimasukkan pengguna
	hitungGaji(nama, gajipokok, jamLembur)
}