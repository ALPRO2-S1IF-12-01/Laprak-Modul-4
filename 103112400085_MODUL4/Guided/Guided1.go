// Anastasia Adinda N.I
// Prosedur Gaji
package main

import "fmt"

func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("Nama Karyawan :", nama)
	fmt.Printf("Gaji Pokok : Rp%.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur : Rp%.2f (%d jam x Rp.50,000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji : Rp%.2f\n", totalGaji)

}

func main() {
	var nama string
	var gajiPokok float64
	var jamLembur int

	//Input pengguna
	fmt.Print("Masukan Nama Karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukan Gaji Pokok: ")
	fmt.Scanln(&gajiPokok)

	fmt.Print("Masukan Jumlah Jam Lembur: ")
	fmt.Scanln(&jamLembur)

	//Memanggil prosedur dengan data dari pengguna
	hitungGaji(nama, gajiPokok, jamLembur)
}
