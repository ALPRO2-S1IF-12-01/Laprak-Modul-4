// Muahamad faza fahri aziz || 103112400072
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Print(permutasi(a, c), " ", combination(a, c), "\n")
	fmt.Print(permutasi(b, d), " ", combination(b, d), "\n")

}

func factorial(n int) int {
	var hasil int
	hasil = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / (factorial(r) * (factorial(n - r)))
}
