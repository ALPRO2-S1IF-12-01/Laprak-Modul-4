package main

import (
	"fmt"
)

func cetakFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	cetakFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	fmt.Printf("Faktor dari %d: ", n)
	cetakFaktor(n, 1)
	fmt.Println()
}
