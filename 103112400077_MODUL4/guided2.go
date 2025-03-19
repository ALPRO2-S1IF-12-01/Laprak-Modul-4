package main

import "fmt"

func hitungRataRata(nama string, n1, n2, n3 float64) {

	ratarata := (n1 + n2 + n3) / 3
	status := "Tidak Lulus"

	if ratarata >= 60 {
		status = "Lulus"
	}

	fmt.Println("\n=== Hasil Akademik ===")
	fmt.Println("Nama Mahasiswa :", nama)
	fmt.Printf("Nilai 1 : %.2f\n", n1)
	fmt.Printf("Nilai 2 : %.2f\n", n2)
	fmt.Printf("Nilai 3 : %.2f\n", n3)
	fmt.Printf("Rata-rata : %.2f\n", ratarata)
	fmt.Println("Status:", status)
}

func main() {
	var nama string
	var n1, n2, n3 float64

	fmt.Print("Masukkan Nama Mahasiswa: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan Nilai 1: ")
	fmt.Scan(&n1)

	fmt.Print("Masukkan Nilai 2: ")
	fmt.Scan(&n2)

	fmt.Print("Masukkan Nilai 3: ")
	fmt.Scan(&n3)

	hitungRataRata(nama, n1, n2, n3)
}