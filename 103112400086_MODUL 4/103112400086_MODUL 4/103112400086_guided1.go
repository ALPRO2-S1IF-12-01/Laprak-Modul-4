package main

import "fmt"

//prosedur untk menghitung total gaji karyawan
func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("\n=== Slip Gaji ===")
	fmt.Println("Nama Karyawan : ", nama)
	fmt.Printf("Gaji Pokok : Rp %.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur : Rp %.2f (%d jam x Rp 50.000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji : Rp %.2f\n", totalGaji)
}

func main() { // main redeclared in this block
	var nama string
	var gajiPokok float64
	var jamLembur int

	fmt.Print("Nama : ")
	fmt.Scanln(&nama)

	fmt.Print("Gaji Pokok : ")
	fmt.Scanln(&gajiPokok)

	fmt.Print("Jumlah Jam Lembur : ")
	fmt.Scanln(&jamLembur)

	// memanggil prosedur dengan data dr inputan

	hitungGaji(nama, gajiPokok, jamLembur)

}
