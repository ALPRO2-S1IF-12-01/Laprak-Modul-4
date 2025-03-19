package main

import "fmt"

func hitungGaji(nama string, gajiPokok float64, jamLembur int) {
	bonus := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonus

	fmt.Println("Nama Karyawan:", nama)
	fmt.Printf("Gaji Pokok: Rp%.2f\n", gajiPokok)
	fmt.Printf("bonus: Rp%.2f(%d jam x Rp50000)\n", bonus, jamLembur)
	fmt.Printf("Total Gaji: Rp%.2f\n", totalGaji)
}

func main() {
	var nama string
	var gajiPokok float64
	var jamLembur int

	fmt.Print("Nama Karyawan: ")	
	fmt.Scanln(&nama)

	fmt.Print("Gaji Pokok: Rp")
	fmt.Scanln(&gajiPokok)

	fmt.Print("Jam Lembur: ")
	fmt.Scanln(&jamLembur)

	hitungGaji(nama, gajiPokok, jamLembur)
}