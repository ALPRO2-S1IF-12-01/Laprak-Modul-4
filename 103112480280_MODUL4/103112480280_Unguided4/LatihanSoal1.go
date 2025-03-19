//Nama : Anggun Wahyu W. (103112480280)
package main

import "fmt"

// Prosedur untuk menghitung faktorial
func factorial(n int, hasil *int) {
    *hasil = 1
    for i := 1; i <= n; i++ {
        *hasil *= i
    }
}

// Prosedur untuk menghitung permutasi
func permutation(n int, r int, hasil *int) {
    var faktorialN, faktorialNR int
    factorial(n, &faktorialN)
    factorial(n-r, &faktorialNR)
    *hasil = faktorialN / faktorialNR
}

// Prosedur untuk menghitung kombinasi
func combination(n int, r int, hasil *int) {
    var faktorialN, faktorialR, faktorialNR int
    factorial(n, &faktorialN)
    factorial(r, &faktorialR)
    factorial(n-r, &faktorialNR)
    *hasil = faktorialN / (faktorialR * faktorialNR)
}

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    if a < c || b < d {
        fmt.Println("Input tidak valid! Pastikan a >= c dan b >= d.")
        return
    }

    // Menghitung dan mencetak hasil permutasi dan kombinasi
    var hasilPermutasiAC, hasilKombinasiAC int
    var hasilPermutasiBD, hasilKombinasiBD int

    permutation(a, c, &hasilPermutasiAC)
    combination(a, c, &hasilKombinasiAC)

    permutation(b, d, &hasilPermutasiBD)
    combination(b, d, &hasilKombinasiBD)

    fmt.Println(hasilPermutasiAC, hasilKombinasiAC)
    fmt.Println(hasilPermutasiBD, hasilKombinasiBD)
}