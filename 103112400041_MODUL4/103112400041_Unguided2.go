//BERTHA ADELA
//103112400041
package main
import "fmt"
func main() {
    var nama, pemenang string
    var skorTertinggi, total, soal, skorPeserta, soalTerjawab int
    soal = 8
    selesai := false
    for selesai == false {
        fmt.Scan(&nama)
        if nama == "Selesai" {
            selesai = true
        } else {
            skorPeserta, soalTerjawab = hitungSkor(soal)
            total += skorPeserta
            if skorPeserta > skorTertinggi {
                skorTertinggi = skorPeserta
                pemenang = nama
            }
        }
    }
    if pemenang != "" {
        fmt.Print(pemenang, " ")
        fmt.Print(soalTerjawab, " ")
        fmt.Print(skorTertinggi, " ")
    }
}

//BERTHA ADELA
//103112400041
func hitungSkor(soal int) (int, int) {
    var totalSkor, totalSoal, skor int
    soal = 8
    for i := 1; i <= soal; i++ {
        fmt.Scan(&skor)
        if skor == 301 {
            totalSkor -= 301
            totalSoal -= 1
        }
        totalSkor += skor
        totalSoal++
    }
    return totalSkor, totalSoal
}