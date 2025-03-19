//RAJA MUHAMMAD LUFHTI 103112400027
package main

import "fmt"

func hitungRataRata(nama string, nilai1, nilai2, nilai3 float64) {
	rataRata := (nilai1 + nilai2 + nilai3) / 3
	status := "Tidak Lulus"
	if rataRata >= 60 {
		status = "Lulus"
	}

	fmt.Println("\n=== hasil akademik ===")
	fmt.Println("Nama mahasiswa:", nama)
	fmt.Printf("Nilai 1: %.2f\n", nilai1)
	fmt.Printf("Nilai 2: %.2f\n", nilai2)
	fmt.Printf("Nilai 3: %.2f\n", nilai3)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)
	fmt.Println("Status:", status)
}

func main() {
	var nama string
	var nilai1, nilai2, nilai3 float64

	fmt.Print("Nama mahasiswa: ")
	fmt.Scanln(&nama)

	fmt.Print("Nilai 1: ")
	fmt.Scanln(&nilai1)

	fmt.Print("Nilai 2: ")
	fmt.Scanln(&nilai2)

	fmt.Print("Nilai 3: ")
	fmt.Scanln(&nilai3)

	hitungRataRata(nama, nilai1, nilai2, nilai3)
}