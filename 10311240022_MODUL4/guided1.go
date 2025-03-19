package main 
import "fmt"

func hitungGaji(nama string, gajiPokok float64, jamLembur int){
	bonusLembur := float64(jamLembur) * 50000
	totalGaji := gajiPokok + bonusLembur

	fmt.Println("n=== Slip Gaji ===")
	fmt.Println("Nama karyawan :", nama)
	fmt.Printf("Gaji Pokok		: Rp%.2f\n", gajiPokok)
	fmt.Printf("Bonus Lembur	: Rp%.2f (%d jam x 50,000)\n", bonusLembur, jamLembur)
	fmt.Printf("Total Gaji		: Rp%.2f\n", totalGaji)	
}

func main(){
	var nama string
	var gajiPokok float64
	var jamLembur int

	fmt.Print("Masukan nama karyawan: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukan gaji pokok: ")
	fmt.Scanln(&gajiPokok)

	fmt.Print("Masukan jumlah jam lembur: ")
	fmt.Scanln(&jamLembur)

	hitungGaji(nama, gajiPokok, jamLembur)
}