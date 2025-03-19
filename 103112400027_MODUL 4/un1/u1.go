//RAJA MUHAMMAD LUFHTI 103112400027
package main

import (
	"fmt" 
)

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan empat angka (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println("Hasil:")
	fmt.Printf("P(%d, %d) = %d\n", a, c, permutation(a, c))
	fmt.Printf("C(%d, %d) = %d\n", a, c, combination(a, c))
	fmt.Printf("P(%d, %d) = %d\n", b, d, permutation(b, d))
	fmt.Printf("C(%d, %d) = %d\n", b, d, combination(b, d))
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

