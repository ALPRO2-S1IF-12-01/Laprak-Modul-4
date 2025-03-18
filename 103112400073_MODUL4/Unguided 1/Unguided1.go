//Muhammad Zaky Mubarok
package main

import "fmt"


func hitungFaktorial(n int, hasil *int) {
	if n == 0 || n == 1 {
		*hasil = 1
		return
	}
	
	var temp int
	hitungFaktorial(n-1, &temp)
	*hasil = n * temp
}


func hitungPermutasi(n, r int, hasil *int) {
	var faktorialN, faktorialNR int
	hitungFaktorial(n, &faktorialN)
	hitungFaktorial(n-r, &faktorialNR)
	*hasil = faktorialN / faktorialNR
}


func hitungKombinasi(n, r int, hasil *int) {
	var faktorialN, faktorialR, faktorialNR int
	hitungFaktorial(n, &faktorialN)
	hitungFaktorial(r, &faktorialR)
	hitungFaktorial(n-r, &faktorialNR)
	*hasil = faktorialN / (faktorialR * faktorialNR)
}

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)

	if a >= c && b >= d {
		
		var p1, c1, p2, c2 int
		hitungPermutasi(a, c, &p1)
		hitungKombinasi(a, c, &c1)
		hitungPermutasi(b, d, &p2)
		hitungKombinasi(b, d, &c2)

		
		fmt.Printf("%d %d\n", p1, c1)
		fmt.Printf("%d %d\n", p2, c2)
	} else {
		fmt.Println("Input tidak valid: a harus lebih besar atau sama dengan c, dan b harus lebih besar atau sama dengan d.")
	}
}
