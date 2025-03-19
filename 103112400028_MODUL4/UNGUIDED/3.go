// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main
import "fmt"
func cetakDeret(n int) {
    fmt.Print(n)
        for n != 1 {
            if n%2 == 0 { 
                n = n / 2
            } else { 
                n = 3*n + 1
            }
            fmt.Print(" ", n)
        }
        fmt.Println() 
}

func main() {
        var masukan int
        fmt.Print("masukkan bilngan positif: ")
        fmt.Scan(&masukan)

        cetakDeret(masukan)
}
