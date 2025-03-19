//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func faktorial(n int, hasil *int) {
    *hasil = 1
    i := 1
    for i <= n {
        *hasil *= i
        i++
    }
}

func permutasi(n, r int, hasil *int) {
    var pembilang, penyebut int
    faktorial(n, &pembilang)
    faktorial(n-r, &penyebut)
    *hasil = pembilang / penyebut
}

func kombinasi(n, r int, hasil *int) {
    var pembilang, penyebut1, penyebut2 int
    faktorial(n, &pembilang)
    faktorial(r, &penyebut1)
    faktorial(n-r, &penyebut2)
    *hasil = pembilang / (penyebut1 * penyebut2)
}

func hitungHasil(n, r int) (p, c int) {
    permutasi(n, r, &p)
    kombinasi(n, r, &c)
    return
}

func main() {
    var a, b, c, d int
    fmt.Print("Masukkan empat bilangan (a b c d): ")
    fmt.Scan(&a, &b, &c, &d)

    p1, c1 := hitungHasil(a, c)
    p2, c2 := hitungHasil(b, d)
    fmt.Printf("%d %d\n", p1, c1)
    fmt.Printf("%d %d\n", p2, c2)
}