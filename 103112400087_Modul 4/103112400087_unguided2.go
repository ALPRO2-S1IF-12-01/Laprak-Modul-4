// 103112400087 Muhammad Shabrian Fadly
package main

import (
	"fmt"
)

func hitungSkor(waktuPengerjaan [8]int, totalSoal *int, totalSkor *int) {
	*totalSoal = 0
	*totalSkor = 0
	for _, waktu := range waktuPengerjaan {
		if waktu <= 300 {
			*totalSoal++
			*totalSkor += waktu
		}
	}
}
func main() {
	var namaPemenang string
	var maxSoal, minSkor int
	for {
		var nama string
		fmt.Print("Masukkan nama peserta (atau ketik 'Selesai' untuk mengakhiri): ")
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		var waktuPengerjaan [8]int
		fmt.Print("Masukkan waktu pengerjaan (8 integer dipisahkan spasi): ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktuPengerjaan[i])
		}
		var totalSoal, totalSkor int
		hitungSkor(waktuPengerjaan, &totalSoal, &totalSkor)
		if totalSoal > maxSoal || (totalSoal == maxSoal && totalSkor < minSkor) {
			namaPemenang = nama
			maxSoal = totalSoal
			minSkor = totalSkor
		}
	}
	fmt.Printf("%s %d %d\n", namaPemenang, maxSoal, minSkor)
}
