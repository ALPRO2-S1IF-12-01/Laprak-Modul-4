package main

import "fmt"

func cetakDeret(n int) {
	if n <= 0 {
		fmt.Println("Error: Bilangan harus positif!")
		return
	}
	langkah := 1
	fmt.Printf("Deret Collatz untuk n = %d:\n", n)
	fmt.Print(n)
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(" → ", n)
		langkah++
	}
	fmt.Println()
	fmt.Printf("Jumlah langkah: %d\n", langkah)
}

func main() {
	fmt.Println("Program Deret Collatz (3n+1)")
	fmt.Println("===========================")

	var masukan int
	fmt.Print("Masukkan bilangan positif: ")
	_, err := fmt.Scan(&masukan)

	if err != nil {
		fmt.Println("Error: Input tidak valid!")
		return
	}

	cetakDeret(masukan)
}
