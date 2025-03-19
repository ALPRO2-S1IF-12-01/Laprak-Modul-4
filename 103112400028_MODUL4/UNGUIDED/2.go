// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func main() {
    var pemenang string
    maxSoal, minSkor := 0, 1<<31-1

    for {
        var nama string
        waktu := make([]int, 8)
        fmt.Scan(&nama)
        if nama == "Selesai" {
            break
        }
        for i := range waktu {
            fmt.Scan(&waktu[i])
        }

        soal, skor := hitungSkor(waktu)
        if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
            pemenang, maxSoal, minSkor = nama, soal, skor
        }
    }
	
    fmt.Printf("Pemenang: %s, Soal yang diselesaikan: %d, Total waktu: %d menit\n", pemenang, maxSoal, minSkor)
}
func hitungSkor(waktu []int) (soal, skor int) {
    for _, w := range waktu {
        if w < 301 {
            soal++
            skor += w
        }
    }
    return
}