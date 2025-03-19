//ABID FADHILAH M
//103112400046
package main

import "fmt"

func hitungSkor(w int, tW, tS *int) {
	if w <= 300 {
		*tS++
		*tW += w
	}
}

func main() {
	var p, p2, n string
	var w, tW1, tW2, tS1, tS2 int

	fmt.Scan(&p)
	if p == "Selesai" {
		return
	}

	for i := 0; i < 8; i++ {
		fmt.Scan(&w)
		hitungSkor(w, &tW1, &tS1)
	}

	n, mS, mW := p, tS1, tW1

	for {
		fmt.Scan(&p2)
		if p2 == "Selesai" {
			break
		}

		tW2, tS2 = 0, 0
		for i := 0; i < 8; i++ {
			fmt.Scan(&w)
			hitungSkor(w, &tW2, &tS2)
		}

		if tS2 > mS || (tS2 == mS && tW2 < mW) {
			n, mS, mW = p2, tS2, tW2
		}
	}

	fmt.Println(n, mS, mW)
}
