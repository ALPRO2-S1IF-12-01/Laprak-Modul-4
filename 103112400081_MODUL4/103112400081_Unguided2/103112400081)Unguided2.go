//RYANAKEYLANOVIANTO
//103112400081

package main

import "fmt"

const maxTime = 301

type Peserta struct {
	nama             string
	waktuSoal        []int
	soalDiselesaikan int
	totalWaktu       int
}

func main() {
	var pemenang Peserta
	for {
		p := inputPeserta()
		if p.nama == "Selesai" {
			break
		}
		hitungSkor(&p)
		if p.soalDiselesaikan > pemenang.soalDiselesaikan || (p.soalDiselesaikan == pemenang.soalDiselesaikan && p.totalWaktu < pemenang.totalWaktu) {
			pemenang = p
		}
	}
	fmt.Printf("Pemenang: %s, Soal diselesaikan: %d, Total waktu: %d menit\n", pemenang.nama, pemenang.soalDiselesaikan, pemenang.totalWaktu)

}

func inputPeserta() Peserta {
	var nama string
	var waktuSoal []int
	fmt.Scanln(&nama)
	if nama == "Selesai" {
		return Peserta{nama: nama}
	}
	waktuSoal = make([]int, 8)
	for i := 0; i < 8; i++ {
		fmt.Scanln(&waktuSoal[i])
	}
	return Peserta{nama: nama, waktuSoal: waktuSoal}
}

func hitungSkor(p *Peserta) {
	p.soalDiselesaikan = 0
	p.totalWaktu = 0
	for _, waktu := range p.waktuSoal {
		if waktu < maxTime {
			p.soalDiselesaikan++
			p.totalWaktu += waktu
		}
	}
}
