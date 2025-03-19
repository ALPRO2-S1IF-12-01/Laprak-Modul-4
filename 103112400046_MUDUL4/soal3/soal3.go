//ABID FADHILAH M
//103112400046
package main

import "fmt"

func cetak(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n int
	fmt.Scan(&n)
	if n < 1000000 {
		cetak(n)
	} else {
		fmt.Println("Eror")
	}
}
