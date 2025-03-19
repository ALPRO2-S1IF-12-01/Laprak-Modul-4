// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n % 2 == 0 {
			n /= 2
		} else {
			n = (3 * n) + 1
		}
	}
	fmt.Print(n)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}