//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func hitungSkor(jumlahSoal *int, totalWaktu *int, namaPemenang *string) {
    var namaPeserta string
    var waktuPengerjaan int

    *jumlahSoal = 0
    *totalWaktu = 0
    *namaPemenang = ""

    for {
        fmt.Scan(&namaPeserta)
        if namaPeserta == "Selesai" {
            return
        }
        soalBerhasil := 0
        waktuBerhasil := 0

        for i := 0; i < 8; i++ {
            fmt.Scan(&waktuPengerjaan)
            if waktuPengerjaan <= 300 {
                soalBerhasil++
                waktuBerhasil += waktuPengerjaan
            }
        }
        if soalBerhasil > *jumlahSoal || 
           (soalBerhasil == *jumlahSoal && waktuBerhasil < *totalWaktu) {
            *jumlahSoal = soalBerhasil
            *totalWaktu = waktuBerhasil
            *namaPemenang = namaPeserta
        }
    }
}

func main() {
    var soalTerbaik, waktuTerbaik int
    var namaPemenang string
    
    hitungSkor(&soalTerbaik, &waktuTerbaik, &namaPemenang)
    fmt.Printf("%s %d %d\n", namaPemenang, soalTerbaik, waktuTerbaik)
}