// MULIA AKBAR NANDA PRATAMA
// 103112400034
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

        fmt.Print("masukkan empat bilangan: ")
        fmt.Scan(&a, &b, &c, &d)

        permAC := permutasi(a, c)
        combAC := kombinasi(a, c)
        fmt.Printf("%d %d\n", permAC, combAC)

        permBD := permutasi(b, d)
        combBD := kombinasi(b, d)
        fmt.Printf("%d %d\n", permBD, combBD)
}