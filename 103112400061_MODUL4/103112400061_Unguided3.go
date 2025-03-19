// KEISHIN NAUFA ALFARIDZHI
// 103112400061
package main

import "fmt"

func main() {
	var (
		n int
	)

	fmt.Scan(&n)

	if n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Input Error")
	}
}

func cetakDeret(n int) {
	var deret int
	for {
		if n%2 == 0 {
			n = n/2
			if n == 1 {
				break
			}
			deret = n
		} else {
			n = 3*n + 1
			deret = n
		}
		fmt.Printf("%v ", deret)
	}
	

	fmt.Println(n)
}

