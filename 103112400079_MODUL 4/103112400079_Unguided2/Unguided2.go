package main

// Muhammad Faris Rachmadi || 103112400079
import "fmt"

func hitungSkor(waktu []int, jumlahSoal *int, totalWaktu *int) {
	batasWaktu := 301
	*jumlahSoal, *totalWaktu = 0, 0
	for _, t := range waktu {
		if t < batasWaktu {
			*jumlahSoal++
			*totalWaktu += t
		}
	}
}

func main() {
	var pemenang string
	var maxSoal, minWaktu int = 0, 1000000

	for {
		var nama string
		var waktu [8]int
		_, err := fmt.Scan(&nama)
		if err != nil || nama == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			_, err = fmt.Scan(&waktu[i])
			if err != nil {
				break
			}
		}

		var jumlahSoal, totalWaktu int
		hitungSkor(waktu[:], &jumlahSoal, &totalWaktu)

		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minWaktu) {
			maxSoal = jumlahSoal
			minWaktu = totalWaktu
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
