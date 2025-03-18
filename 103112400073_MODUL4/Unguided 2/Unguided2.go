//Muhammad Zaky Mubarok
package main

import "fmt"


func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, t := range waktu {
		if t < 301 {
			*soal++
			*skor += t
		}
	}
}


func cariPemenang(peserta map[string][]int, pemenang *string, maxSoal *int, minSkor *int) {
	*pemenang = ""
	*maxSoal = 0
	*minSkor = 9999999

	for nama, waktu := range peserta {
		var soal, skor int
		hitungSkor(waktu, &soal, &skor)
		if soal > *maxSoal || (soal == *maxSoal && skor < *minSkor) {
			*pemenang = nama
			*maxSoal = soal
			*minSkor = skor
		}
	}
}

func main() {
	
	peserta := make(map[string][]int)

	
	fmt.Println("Format: Nama Waktu")
	fmt.Println("Ketik 'Selesai' untuk mengakhiri input. (S nya kapital ya)")

	
	for {
		var nama string
		var waktu [8]int
		fmt.Print("Input: ")

		_, err := fmt.Scan(&nama)
		if err != nil {
			fmt.Println("Input error.")
			break
		}

		if nama == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			_, err := fmt.Scan(&waktu[i])
			if err != nil {
				fmt.Println("Input error.")
				break
			}
		}

		peserta[nama] = waktu[:]
	}

	
	var pemenang string
	var jumlahSoal, totalSkor int
	cariPemenang(peserta, &pemenang, &jumlahSoal, &totalSkor)

	
	fmt.Printf("\n%s %d %d\n", pemenang, jumlahSoal, totalSkor)
}
