// KEISHIN NAUFA ALFARIDZHI
// 103112400061
package main

import "fmt"

var hasilFactorial, hasilPermutation, hasilCombination int

func main() {
	var a, b, c, d, hasil1, hasil2 int
	
	fmt.Scanln(&a, &b, &c, &d)

	permutation(a, c)
	hasil1 = hasilPermutation

	combination(a,c)
	hasil2 = hasilCombination

	fmt.Println(hasil1, hasil2)

	permutation(b, d)
	hasil1 = hasilPermutation

	combination(b, d)
	hasil2 = hasilCombination

	fmt.Println(hasil1, hasil2)
}

func factorial(n int) {
	hasilFactorial = 1
	for i := 1; i <= n ; i++ {
		hasilFactorial *= i
	}
}

func permutation(n, r int) {
	if r>n {
		hasilPermutation = 0
		return
	}

	factorial(n)
	nFactorial := hasilFactorial

	factorial(n-r)
	nrFactorial := hasilFactorial

	hasilPermutation = nFactorial / nrFactorial
}

func combination(n, r int) {
	if r>n {
		hasilCombination = 0
		return
	}

	factorial(n)
	nFactorial := hasilFactorial

	factorial(r)
	rFactorial := hasilFactorial

	factorial(n-r)
	nrFactorial := hasilFactorial

	hasilCombination = nFactorial / (rFactorial * nrFactorial)
}