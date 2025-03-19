package main

import "fmt"

//ARIEL AHNAF KUSUMA 103112400050
func hitungSkor(waktu int, totalWaktu, totalSoal *int) {
	if waktu <= 300 {
		*totalSoal++
		*totalWaktu += waktu
	}
}

func main() {
	var namaPemenang, peserta string
	var waktu, maxSoal, minWaktu int
	var totalWaktu, totalSoal int

	fmt.Scan(&namaPemenang)
	if namaPemenang == "Selesai" {
		return
	}

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		hitungSkor(waktu, &totalWaktu, &totalSoal)
	}

	maxSoal, minWaktu = totalSoal, totalWaktu

	for {
		fmt.Scan(&peserta)
		if peserta == "Selesai" {
			break
		}

		pesertaWaktu, pesertaSoal := 0, 0

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu)
			hitungSkor(waktu, &pesertaWaktu, &pesertaSoal)
		}

		if pesertaSoal > maxSoal || (pesertaSoal == maxSoal && pesertaWaktu < minWaktu) {
			namaPemenang, maxSoal, minWaktu = peserta, pesertaSoal, pesertaWaktu
		}
	}

	fmt.Println(namaPemenang, maxSoal, minWaktu)
}
