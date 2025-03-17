//Anastasia Adinda
//Rekursif

package main

import "fmt"

func pangkatRekursif(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * pangkatRekursif(base, exp-1)
}

func faktorialRekrusif(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorialRekrusif(n-1)
}

func main() {
	var base, exp, n int

	fmt.Print("Masukan Bilangan: ")
	fmt.Scanln(&base)
	fmt.Print("Masukan Pangkat: ")
	fmt.Scanln(&exp)

	fmt.Printf("%d^%d = %d\n", base, exp, pangkatRekursif(base, exp))

	fmt.Print("Masukan angka untuk faktorial: ")
	fmt.Scanln(&n)

	fmt.Printf("%d! = %d\n", n, faktorialRekrusif(n))
}
