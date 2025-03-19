//Nama : Anggun Wahyu W. (103112480280)
package main

import "fmt"

// Prosedur untuk mencetak deret berdasarkan nilai awal
func cetakDeret(n int) {
    for n != 1 {
        fmt.Print(n, " ")
        if n%2 == 0 { 
            n = n / 2
        } else { 
            n = 3*n + 1
        }
    }
    fmt.Print(n) // Mencetak suku terakhir yaitu 1
}

func main() {
    var n int

    fmt.Print("Masukkan bilangan bulat positif (kurang dari 1000000): ")
    fmt.Scan(&n)

    if n <= 0 || n >= 1000000 {
        fmt.Println("Input tidak valid! Pastikan bilangan positif dan kurang dari 1000000.")
        return
    }

    // Memanggil prosedur untuk mencetak deret
    cetakDeret(n)
}