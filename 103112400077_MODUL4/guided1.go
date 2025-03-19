package main

import "fmt"

func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("Nama karyawan:", nama)
	fmt.Printf("Gaji pokok: Rp.%2f\n", gajiPokok)
	fmt.Printf("Bonus lembur: Rp.%2f (%d jam x Rp.50,000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total gaji: Rp.%2f\n", totalGaji)
}

func main() {
	var nama string
	var gajiPokok float64
	var jamLembur int

	fmt.Print("Masukkan nama karyawan:")
	fmt.Scan(&nama)

	fmt.Print("Masukkan gaji pokok:")
	fmt.Scan(&gajiPokok)

	fmt.Print("Masukkan jam lembur:")
	fmt.Scan(&jamLembur)

	hitungGaji(nama, gajiPokok, jamLembur)

}
