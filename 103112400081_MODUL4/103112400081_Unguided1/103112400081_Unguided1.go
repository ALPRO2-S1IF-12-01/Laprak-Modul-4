//RYANAKEYLANOVIANTO
//103112400081

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println(permutasi(a, c), kombinasi(a, c))
		fmt.Println(permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println(permutasi(c, a), kombinasi(c, a))
		fmt.Println(permutasi(d, b), kombinasi(d, b))
	}
}

func kombinasi(n, r int) *big.Int {
	if r > n || r < 0 {
		return big.NewInt(0)
	}
	return new(big.Int).Div(faktorial(n), new(big.Int).Mul(faktorial(r), faktorial(n-r)))
}

func permutasi(n, r int) *big.Int {
	if r > n || r < 0 {
		return big.NewInt(0)
	}
	return new(big.Int).Div(faktorial(n), faktorial(n-r))
}

func faktorial(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
