//RYANAKEYLANOVIANTO
//103112400081

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeretWhile(n)
	fmt.Println()
}

func cetakDeretWhile(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(1)
}
