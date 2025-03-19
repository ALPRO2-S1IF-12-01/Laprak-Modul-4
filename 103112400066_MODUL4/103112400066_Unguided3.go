// DWI OKTA SURYANINGRUM
package main

import "fmt"

// Prosedur cetakDeret mencetak deret bilangan sesuai aturan Skiena dan Revilla.
// Parameter:
// - n: Bilangan awal untuk memulai deret.
func cetakDeret(n int) {
	// Loop akan terus berjalan selama n tidak sama dengan 1
	for n != 1 {
		fmt.Print(n, " ") // Cetak nilai n saat ini diikuti spasi
		if n%2 == 0 {     // Jika n genap
			n = n / 2 // Hitung suku berikutnya: n = n / 2
		} else { // Jika n ganjil
			n = 3*n + 1 // Hitung suku berikutnya: n = 3n + 1
		}
	}
	fmt.Print(n) // Cetak nilai terakhir (1)
}

func main() {
	var n int
	// Minta pengguna memasukkan bilangan awal
	fmt.Scan(&n)

	// Panggil prosedur cetakDeret untuk mencetak deret
	cetakDeret(n)
}