package main

import "fmt"

func hitunggaji(nama string, gajipokok float64, jamlembur int) {
	bonusLembur := float64(jamlembur) * 5000
	totalgaji := gajipokok + bonusLembur

	fmt.Println("\n=== Slip gaji ===")
	fmt.Println("nama karyawan:", nama)
	fmt.Printf("Gaji pokok:Rp. %2f\n", gajipokok)
	fmt.Printf("Bonus lembur: Rp.%2f(%d jam x Rp.50,000)\n", bonusLembur, jamlembur)
	fmt.Printf("total gaji: Rp. %2f\n", totalgaji)
}

func main() {
	var nama string
	var gajipokok float64
	var jamlembur int

	fmt.Print("Masukkan nama karyawan:")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan gaji Pokok:")
	fmt.Scanln(&gajipokok)
	fmt.Print("Masukkan jumlah jam lembur")
	fmt.Scanln(&jamlembur)

	hitunggaji(nama, gajipokok, jamlembur)
}
