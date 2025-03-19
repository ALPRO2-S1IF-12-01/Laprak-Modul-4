package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var nfakt, nrfakt int
	faktorial(n, &nfakt)
	faktorial(n-r, &nrfakt)
	*hasil = nfakt / nrfakt
}

func kombinasi(n, r int, hasil *int) {
	var nfakt, rfakt, nrfakt int
	faktorial(n, &nfakt)
	faktorial(r, &rfakt)
	faktorial(n-r, &nrfakt)
	*hasil = nfakt / (rfakt * nrfakt)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var permA, kombA, permB, kombB int

	permutasi(a, c, &permA)
	kombinasi(a, c, &kombA)

	permutasi(b, d, &permB)
	kombinasi(b, d, &kombB)

	fmt.Println(permA, kombA)
	fmt.Println(permB, kombB)
}
