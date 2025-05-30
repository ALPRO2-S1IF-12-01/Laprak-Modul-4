package main

import (
	"fmt"
)

// Faktorial menghitung n!
func Faktorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Permutasi P(n, r) = n! / (n-r)!
func Permutasi(n, r int) int {
	if r > n {
		return 0
	}
	return Faktorial(n) / Faktorial(n-r)
}

// Kombinasi C(n, r) = n! / (r! * (n-r)!)
func Kombinasi(n, r int) int {
	if r > n {
		return 0
	}
	return Faktorial(n) / (Faktorial(r) * Faktorial(n-r))
}

// Prosedur untuk mencetak hasil permutasi dan kombinasi
func CetakHasil(n, r int) {
	fmt.Println(Permutasi(n, r), Kombinasi(n, r))
}

func main() {
	var a, c, b, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		CetakHasil(a, c)
		CetakHasil(b, d)
	} else {
		fmt.Println("Input tidak valid: pastikan a >= c dan b >= d.")
	}
}
