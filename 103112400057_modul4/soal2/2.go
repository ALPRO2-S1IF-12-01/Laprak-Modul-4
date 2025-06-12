package main

import "fmt"

const maxTime = 301

func hitungSkor(waktu [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, waktuSoal := range waktu {
		if waktuSoal < maxTime {
			*soal++
			*skor += waktuSoal
		}
	}
}

func main() {
	var pemenang string
	var soalPemenang, skorPemenang int
	soalPemenang = -1
	skorPemenang = maxTime * 8
	pesertaCount := 0

	for {
		var nama string
		fmt.Print("Nama peserta (Selesai untuk berhenti): ")
		fmt.Scan(&nama)

		if nama == "Selesai" {
			if pesertaCount == 0 {
				fmt.Println("Belum ada peserta!")
				continue
			}
			break
		}

		var waktu [8]int
		fmt.Print("Waktu (8 soal): ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, skor int
		hitungSkor(waktu, &soal, &skor)
		pesertaCount++

		if soalPemenang == -1 || soal > soalPemenang || (soal == soalPemenang && skor < skorPemenang) {
			pemenang = nama
			soalPemenang = soal
			skorPemenang = skor
		}
	}

	fmt.Printf("\nPemenang: %s\n", pemenang)
	fmt.Printf("Soal selesai: %d\n", soalPemenang)
	fmt.Printf("Total waktu: %d menit\n", skorPemenang)
	fmt.Printf("Jumlah peserta: %d\n", pesertaCount)
}
