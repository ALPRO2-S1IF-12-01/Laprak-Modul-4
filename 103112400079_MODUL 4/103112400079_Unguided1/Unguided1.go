package main

// Muhammad Faris Rachmadi
// 103112400079
import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permut(n, r int) int {
	return fact(n) / fact(n-r)
}

func combi(n, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Println("Masukkan 4 angka:")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permut(a, c)
	c1 := combi(a, c)
	p2 := permut(b, d)
	c2 := combi(b, d)

	fmt.Println(p1, c1, p2, c2)

}
