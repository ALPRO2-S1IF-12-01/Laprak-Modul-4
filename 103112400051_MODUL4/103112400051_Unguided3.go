// PRATAMA BINTANG DANISWARA 103112400051
package main

import "fmt"

func deret(x int) {
    for x != 1 {
        y := x / 2
        if y*2 == x {
            x = y
        } else {
            x = 3*x + 1
        }
        fmt.Print(x, " ")
    }
    fmt.Println()
}

func main() {
    var angka int
    fmt.Scan(&angka)
    deret(angka)
}
