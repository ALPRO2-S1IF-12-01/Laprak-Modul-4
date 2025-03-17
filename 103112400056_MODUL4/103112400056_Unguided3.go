// Felix Pedrosa V

package main

import "fmt"

// Fungsi untuk mencetak deret bilangan
func tampilkanDeret(angka int) {
	for angka != 1 {
		fmt.Printf("%d ", angka) // Cetak nilai angka
		if angka%2 == 0 {        // Jika angka genap
			angka = angka / 2
		} else { // Jika angka ganjil
			angka = 3*angka + 1
		}
	}
	fmt.Println(1) // Cetak nilai akhir 1
}

func main() {
	var nilaiAwal int
	fmt.Scan(&nilaiAwal)  // Ambil input dari pengguna
	tampilkanDeret(nilaiAwal) // Panggil fungsi tampilkanDeret
}