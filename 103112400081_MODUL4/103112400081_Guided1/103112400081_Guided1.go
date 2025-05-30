//RYANAKEYLANOVIANTO
//103112400081

package main

import "fmt"

//Prosedur untuk menghitung gaji karyawan
func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("Nama Karyawan :", nama)
	fmt.Printf("Gaji Pokok : Rp%.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur : Rp%.2f (%d jam x Rp50,000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji : Rp%.2f\n", totalGaji)
}

func main() {
	var nama string
	var gajiPokok float64
	var jamLembur int

	//Input dari pengguna
	fmt.Print("Masukkan Nama Karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Gaji Pokok: ")
	fmt.Scanln(&gajiPokok)

	fmt.Print("Masukkan Jumlah Jam Lembur: ")
	fmt.Scanln(&jamLembur)

	//Memanggil prosedur dengan data dari pengguna
	hitungGaji(nama, gajiPokok, jamLembur)
}
