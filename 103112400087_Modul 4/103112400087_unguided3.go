// 103112400087 Muhammad Shabrian Fadly
package main

import "fmt"

func cetakderet(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("masukkan suku awal (n < 1000000): ")
	fmt.Scan(&n)

	if n >= 1000000 {
		fmt.Println("masukan tidak valid. n harus kurang dari 1000000.")
		return
	}

	cetakderet(n)
}
