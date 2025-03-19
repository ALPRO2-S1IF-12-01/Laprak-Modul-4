// Raihan Adi Arba
// 103112400071
package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
func processValues(a, b, c, d int) {
	if a >= c && b >= d {
		// Menghitung permutasi dan kombinasi
		perm1 := permutation(a, c)
		comb1 := combination(a, c)
		perm2 := permutation(b, d)
		comb2 := combination(b, d)
		//hasil
		fmt.Println(perm1, comb1)
		fmt.Println(perm2, comb2)
	} else {
		fmt.Println("input tidak sesuai")
	}
}
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	processValues(a, b, c, d)
}
