package main

import (
	"fmt"
)

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n /= 2 
		} else {
			n = 3*n + 1 
		}
	}
	fmt.Println(n) 
}

func main() {
	var n int
	fmt.Scan(&n)

	if n < 1 || n >= 1000000 {
		fmt.Println("Bilangan harus di antara 1 dan 999999.")
		return
	}

	cetakDeret(n)
}
