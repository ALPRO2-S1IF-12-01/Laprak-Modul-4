//BERTHA ADELA
//103112400041
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(cetakDeret(n))
}
func cetakDeret(n int) int {
	stop := false
	for stop == false {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n/2
		} else {
			n = 3*n+1
		}
		if n == 1 {
			stop = true
		}
	}
	return n
}