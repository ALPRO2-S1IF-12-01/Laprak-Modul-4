package main
//Mohammad Reyhan Aretha Fatin
//103112400078
import "fmt"

func hitungSkor(waktuSoal int, totalWaktu *int, totalSoal *int) {
	if waktuSoal <= 300 {
		*totalSoal += 1   
		*totalWaktu += waktuSoal  
	}
}

func tentukanPemenang(totalSoal1, totalWaktu1, totalSoal2, totalWaktu2 int, nama1, nama2 string) (string, int, int) {
	if totalSoal1 < totalSoal2 {
		return nama2, totalSoal2, totalWaktu2
	} else if totalSoal1 > totalSoal2 {
		return nama1, totalSoal1, totalWaktu1
	} else {
		if totalWaktu1 > totalWaktu2 {
			return nama1, totalSoal1, totalWaktu1
		}
		return nama2, totalSoal2, totalWaktu2
	}
}

func main() {
	var nama1, nama2 string
	var waktuSoal1, waktuSoal2 int
	var totalWaktu1, totalWaktu2, totalSoal1, totalSoal2 int
	var totalSoalSementara, totalWaktuSementara int
	var namaSementara string
	
	fmt.Scan(&nama1)
	if nama1 != "Selesai" {
	
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktuSoal1)
			hitungSkor(waktuSoal1, &totalWaktu1, &totalSoal1)
		}

		for {
			fmt.Scan(&nama2)
			if nama2 == "Selesai" {
				break
			}
			
			for i := 0; i < 8; i++ {
				fmt.Scan(&waktuSoal2)
				hitungSkor(waktuSoal2, &totalWaktu2, &totalSoal2)
			}

			namaSementara, totalSoalSementara, totalWaktuSementara = tentukanPemenang(totalSoal1, totalWaktu1, totalSoal2, totalWaktu2, nama1, nama2)
		}

		fmt.Println(namaSementara, totalSoalSementara, totalWaktuSementara)
	}
}
