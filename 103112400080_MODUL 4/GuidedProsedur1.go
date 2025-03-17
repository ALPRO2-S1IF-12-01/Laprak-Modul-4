package main

import "fmt"

func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 500
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("Nama Karyawan :", nama)
	fmt.Printf("Gaji Pokok :Rp%.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur :Rp%f (%d jam x Rp50,000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji :Rp%.2f\n", totalGaji)
}

func main() {
	var nama string
	var gajiPokok float64
	var jamLembur int

	fmt.Print("Masukkan Nama Karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Gaji Pokok: ")
	fmt.Scanln(&gajiPokok)

	fmt.Print("Masukkan Jumlah Jam Lembur: ")
	fmt.Scan(&jamLembur)

	hitungGaji(nama, gajiPokok, jamLembur)
}
