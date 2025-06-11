// daffa tsaqifna f
package main

import (
	"fmt"
	"sort"
)

func sortandcount(x [8]int) (int, int) {
	var soal, waktu int
	numSlice := x[:]
	sort.Ints(numSlice)
	for _, num := range numSlice {
		if waktu+num > 301 {
			break
		}
		waktu += num
		soal++
	}
	return waktu, soal
}

func main() {
	var nama string
	var Wname []string
	var pengerjaan [8]int
	var Wtime, Wsoal int
	fmt.Scan(&nama)
	for i := 0; i < 8; i++ {
		fmt.Scan(&pengerjaan[i])
	}
	Wtime, Wsoal = sortandcount(pengerjaan)
	Wname = append(Wname, nama)
	for {
		fmt.Scan(&nama)
		if nama == "selesai" {
			break
		}
		for i := 0; i < 8; i++ {
			fmt.Scan(&pengerjaan[i])
		}
		Ntime, Nsoal := sortandcount(pengerjaan)
		if Nsoal > Wsoal || (Nsoal == Wsoal && Ntime < Wtime) {
			Wsoal = Nsoal
			Wtime = Ntime
			Wname = []string{nama}
		} else if Nsoal == Wsoal && Ntime == Wtime {
			Wname = append(Wname, nama)
		}
	}
	fmt.Printf("%s %d %d", Wname[0], Wsoal, Wtime)
}
