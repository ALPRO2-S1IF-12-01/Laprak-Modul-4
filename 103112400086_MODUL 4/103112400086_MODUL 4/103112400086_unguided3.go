package main

import "fmt"

//  3n+1
func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(n)
}

func main() {
	var n int

	fmt.Scan(&n)

	cetakDeret(n)
}
