package main

//HISYAM NURDIATMOKO 103112400049
import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
