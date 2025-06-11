// daffa tsaqifna f
package main

import "fmt"

func faktorial(x int) int {
	var out int
	out = 1
	for i := 1; i <= x; i++ {
		out *= i
	}
	return out
}
func permutasi(x, y int) int {
	return faktorial(x) / faktorial(x-y)
}
func kombinasi(x, y int) int {
	return faktorial(x) / (faktorial(y) * faktorial(x-y))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Print(permutasi(b, d), kombinasi(b, d))
}
