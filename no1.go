package main

import (
	"fmt"
	"strconv"
)

func main() {
	var jumlahTiket int
	var jenisKursi string
	var member bool

	fmt.Print("Masukkan jumlah tiket: ")
	fmt.Scanln(&jumlahTiket)

	fmt.Print("Masukkan jenis kursi (biasa/VIP): ")
	fmt.Scanln(&jenisKursi)

	fmt.Print("Apakah anda member? (true/false): ")
	fmt.Scanln(&member)

	hargaTiket := 0
	if jenisKursi == "VIP" {
		if member {
			hargaTiket = 60000
		} else {
			hargaTiket = 70000
		}
	} else {
		if member {
			hargaTiket = 40000
		} else {
			hargaTiket = 50000
		}
	}

	totalHarga := hargaTiket * jumlahTiket
	if jumlahTiket > 2 {
		totalHarga = totalHarga - (totalHarga)
	}

	fmt.Println("Total biaya pembelian tiket: Rp " + strconv.Itoa(totalHarga))
}
