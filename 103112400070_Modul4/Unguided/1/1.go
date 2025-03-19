//Achmad Zulvan Nur Hakim 103112400070

package main

import (
	"fmt"
)

func Faktorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func Permutasi(n, r int) int {
	return Faktorial(n) / Faktorial(n-r)
}

func Kombinasi(n, r int) int {
	return Faktorial(n) / (Faktorial(r) * Faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	if a < c || b < d {
		return
	}

	permA := Permutasi(a, c)
	kombA := Kombinasi(a, c)
	permB := Permutasi(b, d)
	kombB := Kombinasi(b, d)

	fmt.Println(permA, kombA)
	fmt.Println(permB, kombB)
}
