// daffa tsaqifna f
package main

import "fmt"

func algorithm(x int) {
	fmt.Print(x, " ")
	if x == 1 {
		return
	}
	if x%2 == 0 {
		algorithm(x / 2)
	} else {
		algorithm(x*3 + 1)
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	algorithm(x)
}
