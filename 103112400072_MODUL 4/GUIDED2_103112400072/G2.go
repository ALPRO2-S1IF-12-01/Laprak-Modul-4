package main

import "fmt"

func hitungRataRata(nama string, nilai1, nilai2, nilai3 float64) {
	rataRata := (nilai1 + nilai2 + nilai3) / 3
	status := "Tidak Lulus"
	if rataRata >= 60 {
		status = "Lulus"
	}

	fmt.Println("Masukan Nama Mahasiswa :", nama)
	fmt.Printf("Nilai 1  : %.2f\n", nilai1)
	fmt.Printf("Nilai 2  : %.2f\n", nilai2)
	fmt.Printf("Nilai 3  : %.2f\n", nilai3)
	fmt.Printf("Rata-rata  :%.2f\n", rataRata)
	fmt.Println("Status :", status)
}

func main() {
	var (
		nama                   string
		nilai1, nilai2, nilai3 float64
	)

	fmt.Print("Masukan Nama Mahasiswa :")
	fmt.Scan(&nama)

	fmt.Print("Masukan nilai 1 :")
	fmt.Scan(&nilai1)

	fmt.Print("Masukan nilai 2 :")
	fmt.Scan(&nilai2)

	fmt.Print("Masukan nilai 3 :")
	fmt.Scan(&nilai3)

	hitungRataRata(nama, nilai1, nilai2, nilai3)
}
