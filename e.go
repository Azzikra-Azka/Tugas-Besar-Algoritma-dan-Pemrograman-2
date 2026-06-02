package main

import "fmt"

const NMAX int = 100

type peserta struct {
	nama   string
	bidang string
	aktif  bool
}

type arrPeserta [NMAX]peserta

func statistik(T arrPeserta, n int) {
	var i int
	var jumlahDesain int
	var jumlahProgramming int
	var jumlahDataScience int
	var total int

	i = 0

	for i < n {

		if T[i].bidang == "Desain" {
			jumlahDesain = jumlahDesain + 1
		}

		if T[i].bidang == "Programming" {
			jumlahProgramming = jumlahProgramming + 1
		}

		if T[i].bidang == "Data Science" {
			jumlahDataScience = jumlahDataScience + 1
		}

		if T[i].aktif == true {
			total = total + 1
		}

		i = i + 1
	}

	fmt.Println("Data Statistik KursusIn")
	fmt.Println("Peserta bidang Desain       :", jumlahDesain)
	fmt.Println("Peserta bidang Programming  :", jumlahProgramming)
	fmt.Println("Peserta bidang Data Science :", jumlahDataScience)
	fmt.Println("Jumlah peserta aktif        :", total)
}

func main() {

	var data arrPeserta
	var n int

	n = 4

	data[0].nama = "Budi"
	data[0].bidang = "Desain"
	data[0].aktif = true

	data[1].nama = "Angga"
	data[1].bidang = "Data Science"
	data[1].aktif = false

	data[2].nama = "Citra"
	data[2].bidang = "Programming"
	data[2].aktif = true

	data[3].nama = "Dina"
	data[3].bidang = "Desain"
	data[3].aktif = true

	statistik(data, n)
}