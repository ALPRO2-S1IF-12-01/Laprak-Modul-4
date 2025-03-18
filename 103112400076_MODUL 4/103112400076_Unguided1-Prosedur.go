package main

import (
    "fmt"
)
func faktorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}
func permutasi(n, r int) int {
    return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
    return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
    if a >= c && b >= d {
        fmt.Printf("P(%d,%d) = %d, C(%d,%d) = %d\n", a, c, permutasi(a, c), a, c, kombinasi(a, c))
        fmt.Printf("P(%d,%d) = %d, C(%d,%d) = %d\n", b, d, permutasi(b, d), b, d, kombinasi(b, d))
    } else {
        fmt.Println("Input tidak valid, pastikan a >= c dan b >= d")
    }
}
