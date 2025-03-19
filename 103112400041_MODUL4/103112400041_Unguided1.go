//BERTHA ADELA
//103112400041
package main
import "fmt"
func main() {
	var a, b, c, d int
	fmt.Scan(&a,&b,&c,&d)

	if a >= c && b >= d{
		fmt.Print(Permutasi(a,c), " ")
		fmt.Println(Kombinasi(a,c))
		fmt.Print(Permutasi(b,d), " ")
		fmt.Println(Kombinasi(b,d))
	} else {
		fmt.Print(Permutasi(c,a), " ")
		fmt.Println(Kombinasi(c,a))
		fmt.Print(Permutasi(d,b), " ")
		fmt.Println(Kombinasi(d,b))
	}
}

//BERTHA ADELA
//103112400041
func Faktorial(n int) int {
	hasil := 1
	for i := 1; i<=n; i++ {
		hasil *= i
	}
	return hasil
}
func Kombinasi(n,r int) int {
	if r > n {
		return 0
	}
	return Faktorial(n) / (Faktorial(r) * Faktorial(n-r))
}
func Permutasi(n,r int) int {
	if r > n {
		return 0
	}
	return Faktorial(n) / Faktorial(n-r)
}