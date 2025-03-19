// PRATAMA BINTANG DANISWARA 103112400051
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}
func mutasi(n, r int) int {
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
func main() {
	var w, x, y, z int
	fmt.Scan(&w, &x, &y, &z)

	mutasi1 := mutasi(w, y)
	kombinasi1 := kombinasi(w, y)
	mutasi2 := mutasi(x, z)
	kombinasi2 := kombinasi(x, z)
	fmt.Printf("%d %d\n", mutasi1, kombinasi1)
	fmt.Printf("%d %d\n", mutasi2, kombinasi2)

}
