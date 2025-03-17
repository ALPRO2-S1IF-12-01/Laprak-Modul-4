//Anastasia Adinda

package main

import (
	"fmt"
)

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}

func kombinasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var hasilP1, hasilC1, hasilP2, hasilC2 int

	fmt.Print("Masukkan empat bilangan asli (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		permutasi(a, c, &hasilP1)
		kombinasi(a, c, &hasilC1)
		permutasi(b, d, &hasilP2)
		kombinasi(b, d, &hasilC2)
		fmt.Printf("%d %d\n", hasilP1, hasilC1)
		fmt.Printf("%d %d\n", hasilP2, hasilC2)
	} else {
		fmt.Println("Input tidak memenuhi syarat a >= c dan b >= d")
	}
}
