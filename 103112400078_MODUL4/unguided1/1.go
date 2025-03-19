package main
// Mohammad Reyhan Aretha Fatin
// 103112400078
import "fmt"

func main() {
	var a, b, c, d int
	var permut, komb int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		permutasi(a, c, &permut)
		kombinasi(a, c, &komb)
		fmt.Println(permut, komb)
	} else {
		permutasi(c, a, &permut)
		kombinasi(c, a, &komb)
		fmt.Println(permut, komb)
	}
	if b >= d {
		permutasi(b, d, &permut)
		kombinasi(b, d, &komb)
		fmt.Println(permut, komb)
	} else {
		permutasi(d, b, &permut)
		kombinasi(d, b, &komb)
		fmt.Println(permut, komb)
	}
}

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
	}
	faktorial(n, &test1)
	faktorial(n-r, &test2)
	*hasil = test1 / test2
}

func kombinasi(n, r int, hasil *int) {
	var test1, test2, test3 int
	if r > n {
		*hasil = 0
	}
	faktorial(n, &test1)
	faktorial(n-r, &test2)
	faktorial(r, &test3)
	*hasil = test1 / (test2 * test3)
}