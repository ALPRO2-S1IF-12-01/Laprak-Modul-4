//RAIHAN ADI ARBA
//103112400071

package main

import "fmt"

func CetakDeret(n int, hasil *float64) {
	fmt.Print(n, " ")

	var ntemp int = n
	for ntemp != 1 {
		if ntemp%2 == 0 {
			*hasil = float64(ntemp) / 2
			ntemp = int(*hasil)
		} else {
			*hasil = 3*float64(ntemp) + 1
			ntemp = int(*hasil)
		}
		fmt.Print(ntemp, " ")
	}
}

func main() {
	var n int
	var deret float64
	fmt.Scan(&n)
	if n < 1000000 {
		CetakDeret(n, &deret)
	} else {
		fmt.Print("Eror")
	}
}
