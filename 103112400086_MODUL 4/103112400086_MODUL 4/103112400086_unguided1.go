package main

import "fmt"

var a, b, c, d int

func faktorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func hitungPermutasiKombinasi(a, b, c, d int) {
	if a < c || b < d {
		fmt.Println("Input tidak memenuhi syarat: a ≥ c dan b ≥ d")
		return
	}
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))

}

func main() {
	fmt.Scan(&a, &b, &c, &d)

	hitungPermutasiKombinasi(a, b, c, d)
}
