package main

import "fmt"

//Prosedur untuk menghitung total gaji karyawan
func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Printf("Nama Karyawan : %s\n", nama)
	fmt.Printf("Gaji Pokok     : Rp%.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur   : Rp%.2f (%d jam x Rp50,000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji     : Rp%.2f\n", totalGaji)
}

func main() {
	var nama string
	var gajiPokok float64
	var jamLembur int

	//Menerima input dari user
	fmt.Print("Masukkan Nama Karyawan : ")
	fmt.Scanln(&nama)
	fmt.Println("Masukkan Gaji Pokok  : ")
	fmt.Scanln(&gajiPokok)
	fmt.Println("Masukkan Jam Lembur  : ")
	fmt.Scanln(&jamLembur)

	//Menghitung total gaji
	hitungGaji(nama, gajiPokok, jamLembur)
}
