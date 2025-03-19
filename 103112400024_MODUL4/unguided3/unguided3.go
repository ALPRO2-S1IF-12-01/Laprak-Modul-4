package main
//103112400024
//SETYO NUGROHO
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
	fmt.Print(n) 
}

func main() {
	var n int
	fmt.Scan(&n)
	if n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Input Tidak Valid")
	}
}
