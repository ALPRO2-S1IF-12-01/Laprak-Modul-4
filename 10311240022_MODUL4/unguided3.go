//Damanik, Yohanes Geovan Ondova
//10311240022

package main 
import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(n, " ")
	}
}

func main(){
	var n int
	fmt.Scan(&n)
	if n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Print("Eror")
	}
}