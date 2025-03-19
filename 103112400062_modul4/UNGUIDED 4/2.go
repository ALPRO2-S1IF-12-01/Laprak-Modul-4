// M.DAVI ILYAS RENALDO
// 103112400062
package main
import "fmt"

func main() {
	var pemenang string
	var totalSoal, skorPemenang int
	totalSoal = 0
	skorPemenang = maxTime * 8

	for {
		var nama string
		var waktu [8]int
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, skor int
		hitungSkor(waktu, &soal, &skor)

		if soal > totalSoal || (soal == totalSoal && skor < skorPemenang) {
			pemenang = nama
			totalSoal = soal
			skorPemenang = skor
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, totalSoal, skorPemenang)
}
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