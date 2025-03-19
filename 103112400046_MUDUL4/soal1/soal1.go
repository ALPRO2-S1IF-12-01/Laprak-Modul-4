//ABID FADHILAH M
//103112400046
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println(p(a, c), k(a, c))
		fmt.Println(p(b, d), k(b, d))
	} else {
		fmt.Println("tidak sesuai")
	}
}

func faktorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func p(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}

func k(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
