package main

import "fmt"

//ARIEL AHNAF KUSUMA 103112400050
func fact(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func perm(n, r int) int {
	return fact(n) / fact(n-r)
}

func comb(n, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(perm(a, c), comb(a, c))
	fmt.Println(perm(b, d), comb(b, d))
}
