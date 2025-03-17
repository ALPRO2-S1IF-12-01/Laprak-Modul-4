//Feros Pedrosa

package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var test1, test2 int
	if r > n {
		*hasil = 0
		return
	}
	faktorial(n, &test1)
	faktorial(n-r, &test2)
	*hasil = test1 / test2
}

func kombinasi(n, r int, hasil *int) {
	var test1, test2, test3 int
	if r > n {
		*hasil = 0
		return
	}
	faktorial(n, &test1)
	faktorial(n-r, &test2)
	faktorial(r, &test3)
	*hasil = test1 / (test2 * test3)
}

func main() {
	var a, b, c, d int
	var permu, kombi int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		permutasi(a, c, &permu)
		kombinasi(a, c, &kombi)
		fmt.Println(permu, kombi)

		permutasi(b, d, &permu)
		kombinasi(b, d, &kombi)
		fmt.Println(permu, kombi)
	} else {
		fmt.Println("input tidak valid")
	}
}
