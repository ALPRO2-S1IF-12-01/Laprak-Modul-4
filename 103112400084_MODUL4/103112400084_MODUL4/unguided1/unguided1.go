//103112400084
//Nufail Alauddin Tsaqif
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / (faktorial(n-r) * faktorial(r))
}

func hitung(n1, n2 int) {
	permut := permutasi(n1, n2)
	komb := kombinasi(n1, n2)
	fmt.Println(permut, komb)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c {
		hitung(a, c)
	} else {
		hitung(c, a)
	}

	if b >= d {
		hitung(b, d)
	} else {
		hitung(d, b)
	}
}
