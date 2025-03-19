//RAJA MUHAMMAD LUFHTI 103112400027
package main

import "fmt"

func hitungNilai(waktu [8]int, soal *int, nilai *int) {
	*soal = 0
	*nilai = 0
	for _, waktuSoal := range waktu {
		if waktuSoal < waktuMaksimal {
			*soal++
			*nilai += waktuSoal
		}
	}
}
const waktuMaksimal = 301

func main() {
	var Nama string
	var nilaiPemenang, soalPemenang int
	soalPemenang = 0
	nilaiPemenang = waktuMaksimal * 8

	for {
		var nama string
		var waktu [8]int
		fmt.Scan(&nama)

		if nama == "selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, nilai int
		hitungNilai(waktu, &soal, &nilai)

		if soal > soalPemenang || (soal == soalPemenang && nilai < nilaiPemenang) {
			soalPemenang = soal
			nilaiPemenang = nilai
			Nama = nama
		}
	}

	fmt.Printf("Pemenang: %s, Soal yang diselesaikan: %d, Total waktu: %d menit\n", Nama, soalPemenang, nilaiPemenang)
}


