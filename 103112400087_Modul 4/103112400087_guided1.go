package main

import "fmt"

func hitunggaji(nama string, gajipokok float64, jamlembur int) {
	bonuslembur := float64(jamlembur) * 50000
	totalgaji := gajipokok + bonuslembur

	fmt.Println("\n=== slip gaji ===")
	fmt.Println("nama karyawan", nama)
	fmt.Printf("gaji pokok  : Rp%.2f\n", gajipokok)
	fmt.Printf("bonus lembur: Rp%.2f\n (%d jam x 50000 )\n", bonuslembur, jamlembur)
	fmt.Printf("total gaji  : Rp%.2f\n", totalgaji)
}

func main() {
	var nama string
	var gajipokok float64
	var jamlembur int

	fmt.Print("masukkan nama karyawan : ")
	fmt.Scanln(&nama)

	fmt.Print("masukkan gaji pokok : ")
	fmt.Scanln(&gajipokok)
	fmt.Print("masukkan jumlah jam lembur:")
	fmt.Scanln(&jamlembur)

	hitunggaji(nama, gajipokok, jamlembur)

}
