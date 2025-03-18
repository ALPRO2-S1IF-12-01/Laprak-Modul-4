//MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"

//prosedur untuk menghitung total gaji karyawan
func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok * bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("Nama Karyawan :", nama)
	fmt.Printf("Gaji Pokok : Rp%.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur : Rp%.2f (%d x Rp50.000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji : Rp%.2f\n", totalGaji)
}

func main() {
	var nama string
	var gajiPokok float64
	var jamLembur int

	//input dari pengguna
	fmt.Print("masukkan nama karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("masukkan gaji pokok: ")
	fmt.Scanln(&gajiPokok)

	fmt.Print("masukkan jumlah jam lembur: ")
	fmt.Scanln(&jamLembur)

	//memangguk proseedur dengan data dari pengguna
	hitungGaji(nama, gajiPokok, jamLembur)
}