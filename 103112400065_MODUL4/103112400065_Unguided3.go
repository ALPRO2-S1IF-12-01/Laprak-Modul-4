// Dimas Ramadhani
// 103112400065

package main

import "fmt"

func cetakDeret(n int, hasil *float64) {
	var ntemp int
	ntemp = n
	for ntemp != 1{
		if ntemp%2==0 {
			*hasil = float64(ntemp)/2
			ntemp = int(*hasil)
		} else if ntemp%2!=0 {
			*hasil = 3*float64(ntemp)+1
			ntemp = int(*hasil)
		}
		fmt.Print(*hasil, " ")
	}
}
func main() {
	var n int
	var deret float64
	fmt.Scan(&n)
	if n < 1000000 {
		cetakDeret(n, &deret)
	} else {
		fmt.Print("Eror")
	}
}	