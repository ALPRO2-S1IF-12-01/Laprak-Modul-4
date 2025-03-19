//DWI OKTA SURYANINGRUM
package main

import "fmt"

// Fungsi untuk menghitung faktorial dari sebuah angka.
// Faktorial adalah hasil perkalian semua angka dari 1 sampai angka tersebut.
func factorial(n int, hasil *int) {
	*hasil = 1 // Mulai dengan hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i // Kalikan hasil dengan angka i (1, 2, 3, ..., n)
	}
}

// Fungsi untuk menghitung permutasi.
// Permutasi adalah jumlah cara mengatur r elemen dari n elemen yang tersedia, dengan memperhatikan urutan.
// Rumus permutasi: P(n, r) = n! / (n-r)!
func permutation(n, r int, hasil *int) {
	var fn, fnr int // Variabel untuk menyimpan faktorial n dan faktorial (n-r)
	factorial(n, &fn) // Hitung faktorial n
	factorial(n-r, &fnr) // Hitung faktorial (n-r)
	*hasil = fn / fnr // Hitung permutasi menggunakan rumus
}

// Fungsi untuk menghitung kombinasi.
// Kombinasi adalah jumlah cara memilih r elemen dari n elemen yang tersedia, tanpa memperhatikan urutan.
// Rumus kombinasi: C(n, r) = n! / (r! * (n-r)!)
func combination(n, r int, hasil *int) {
	var fn, fr, fnr int // Variabel untuk menyimpan faktorial n, r, dan (n-r)
	factorial(n, &fn) // Hitung faktorial n
	factorial(r, &fr) // Hitung faktorial r
	factorial(n-r, &fnr) // Hitung faktorial (n-r)
	*hasil = fn / (fr * fnr) // Hitung kombinasi menggunakan rumus
}

func main() {
	// Variabel untuk menyimpan input dari pengguna
	var a, b, c, d int

	// Variabel untuk menyimpan hasil permutasi dan kombinasi
	var perm1, comb1, perm2, comb2 int

	// Minta pengguna memasukkan 4 angka
	fmt.Scan(&a, &b, &c, &d)

	// Cek apakah input valid (a >= c dan b >= d)
	// Ini penting karena permutasi dan kombinasi hanya bisa dihitung jika n >= r
	if a >= c && b >= d {
		// Hitung permutasi dan kombinasi untuk a dan c
		permutation(a, c, &perm1) // Hitung P(a, c)
		combination(a, c, &comb1) // Hitung C(a, c)

		// Hitung permutasi dan kombinasi untuk b dan d
		permutation(b, d, &perm2) // Hitung P(b, d)
		combination(b, d, &comb2) // Hitung C(b, d)

		// Tampilkan hasil permutasi dan kombinasi untuk a dan c
		fmt.Println(perm1, comb1)

		// Tampilkan hasil permutasi dan kombinasi untuk b dan d
		fmt.Println(perm2, comb2)
	} else {
		// Jika input tidak valid, tampilkan pesan kesalahan
		fmt.Println("input tidak sesuai")
	}
}