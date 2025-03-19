package main

import "fmt"

func main() {
	var (
		nama string
		gajiPokok float64
		jamLembur int
	)

	// Input dari pengguna
	fmt.Print("Masukkan nama karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan gaji pokok: ")
	fmt.Scanln(&gajiPokok)

	fmt.Print("Masukkan jumlah jam lembur: ")
	fmt.Scanln(&jamLembur)

	// Memanggil Prosedur dengan data diri pengguna
	hitungGaji(nama, gajiPokok, jamLembur)
}

func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("Nama Karyawan :", nama)
	fmt.Printf("Gaji Pokok : Rp%.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur : Rp%.2f (%d jam x Rp50,000)\n", gajiPokok, jamLembur)
	fmt.Printf("TotalGaji : Rp%.2f\n", totalGaji)
}