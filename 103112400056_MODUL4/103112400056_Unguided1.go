// Felix Pedrosa V

package main

import "fmt"

func main() {
	var x, y, z, w int
	var permCount, combCount int
	
	fmt.Scan(&x, &y, &z, &w)

	// Menghitung permutasi dan kombinasi untuk pasangan pertama
	if x >= z {
		calculatePermComb(x, z, &permCount, &combCount)
	} else {
		calculatePermComb(z, x, &permCount, &combCount)
	}
	fmt.Println(permCount, combCount)

	// Menghitung permutasi dan kombinasi untuk pasangan kedua
	if y >= w {
		calculatePermComb(y, w, &permCount, &combCount)
	} else {
		calculatePermComb(w, y, &permCount, &combCount)
	}
	fmt.Println(permCount, combCount)
}

func factorial(num int, result *int) {
	*result = 1
	for i := 1; i <= num; i++ {
		*result *= i
	}
}

func calculatePermComb(n, r int, perm *int, comb *int) {
	var factN, factNMinusR, factR int
	if r > n {
		*perm = 0
		*comb = 0
		return
	}
	factorial(n, &factN)
	factorial(n-r, &factNMinusR)
	*perm = factN / factNMinusR

	factorial(r, &factR)
	*comb = factN / (factNMinusR * factR)
}