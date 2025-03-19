//RAJA MUHAMMAD LUFHTI 103112400027
package main

import "fmt"

func cetakDeret(t int) {
	for t != 1 {
		fmt.Print(t, " ")
		if t%2 == 0 {
			t /= 2
		} else {
			t = 3*t + 1
		}
	}
	fmt.Println(1) 
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	cetakDeret(n)		
}