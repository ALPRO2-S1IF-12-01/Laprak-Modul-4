//Ahmad Ruba'i
// 103112400074
package main

import "fmt"

func hitungBerikutnya(bilanganSaatIni int, bilanganBerikut *int) {
    if bilanganSaatIni%2 == 0 {
        *bilanganBerikut = bilanganSaatIni / 2
    } else {
        *bilanganBerikut = 3*bilanganSaatIni + 1
    }
}
func cetakDeret(bilanganAwal int) {
    var nilaiSekarang int = bilanganAwal
    var nilaiBaru int
    fmt.Printf("%d ", nilaiSekarang)

    for nilaiSekarang != 1 {
        hitungBerikutnya(nilaiSekarang, &nilaiBaru)
        nilaiSekarang = nilaiBaru
        fmt.Printf("%d ", nilaiSekarang)
    }
    fmt.Println()
}

func main() {
    var bilanganAwal int
    fmt.Scan(&bilanganAwal)

    if bilanganAwal > 0 && bilanganAwal < 1000000 {
        cetakDeret(bilanganAwal)
    }
}