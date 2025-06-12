package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Println("Program Menghitung Permutasi dan Kombinasi")
	fmt.Println("=========================================")
	fmt.Print("Masukkan empat bilangan (a b c d): ")
	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		fmt.Println("Error: Input tidak valid!")
		return
	}
	if a < 0 || b < 0 || c < 0 || d < 0 {
		fmt.Println("Error: Bilangan tidak boleh negatif!")
		return
	}

	if c > a || d > b {
		fmt.Println("Error: r tidak boleh lebih besar dari n!")
		return
	}
	fmt.Println("\nHasil Perhitungan:")
	fmt.Println("------------------")
	fmt.Printf("Untuk n=%d dan r=%d:\n", a, c)
	fmt.Printf("Permutasi: %d\n", permutasi(a, c))
	fmt.Printf("Kombinasi: %d\n", kombinasi(a, c))
	fmt.Printf("\nUntuk n=%d dan r=%d:\n", b, d)
	fmt.Printf("Permutasi: %d\n", permutasi(b, d))
	fmt.Printf("Kombinasi: %d\n", kombinasi(b, d))
}
