package main

import (
	"fmt"
	"strconv"
)

type Peserta struct {
	ID           int
	Nama         string
	Alamat       string
	BidangMinat  string
	TanggalMasuk string
}

var dataPeserta []Peserta

func binarySearch(nama string) {

	low := 0
	high := len(dataPeserta) - 1

	for low <= high {

		mid := (low + high) / 2

		if dataPeserta[mid].Nama == nama {

			fmt.Println("Data ditemukan")
			fmt.Println(dataPeserta[mid])

			return

		} else if dataPeserta[mid].Nama < nama {

			low = mid + 1

		} else {

			high = mid - 1
		}
	}

	fmt.Println("Data tidak ditemukan")
	return

}

func binarySearchID(id int) {

	low := 0
	high := len(dataPeserta) - 1

	for low <= high {

		mid := (low + high) / 2

		if dataPeserta[mid].ID == id {

			fmt.Println("Data ditemukan")
			fmt.Println(dataPeserta[mid])

			return

		} else if dataPeserta[mid].ID < id {

			low = mid + 1

		} else {

			high = mid - 1
		}
	}

	fmt.Println("Data tidak ditemukan")
	return

}

func insertionSortId() {

	var i, j int
	var temp Peserta

	for i = 1; i < len(dataPeserta); i++ {

		temp = dataPeserta[i]

		j = i - 1

		for j >= 0 && dataPeserta[j].ID > temp.ID {

			dataPeserta[j+1] = dataPeserta[j]

			j--
		}

		dataPeserta[j+1] = temp
	}
}

func insertionSortNama() {

	var i, j int
	var temp Peserta

	for i = 1; i < len(dataPeserta); i++ {

		temp = dataPeserta[i]

		j = i - 1

		for j >= 0 && dataPeserta[j].Nama > temp.Nama {

			dataPeserta[j+1] = dataPeserta[j]

			j--
		}

		dataPeserta[j+1] = temp
	}
	return

}
func selectionSortNama() {

	var i, j, idxMin int
	var temp Peserta

	for i = 0; i < len(dataPeserta)-1; i++ {

		idxMin = i

		for j = i + 1; j < len(dataPeserta); j++ {

			if dataPeserta[idxMin].Nama > dataPeserta[j].Nama {

				idxMin = j
			}
		}

		temp = dataPeserta[idxMin]
		dataPeserta[idxMin] = dataPeserta[i]
		dataPeserta[i] = temp
	}

	return
}

func selectionSort() {

	var i, j, idxMin int
	var temp Peserta

	for i = 0; i < len(dataPeserta)-1; i++ {

		idxMin = i

		for j = i + 1; j < len(dataPeserta); j++ {

			if dataPeserta[idxMin].ID > dataPeserta[j].ID {

				idxMin = j
			}
		}

		temp = dataPeserta[idxMin]
		dataPeserta[idxMin] = dataPeserta[i]
		dataPeserta[i] = temp
	}

	return
}

func statistik() {

	var jumlahCyber int
	var jumlahUIUX int
	var jumlahDataScience int

	for i := 0; i < len(dataPeserta); i++ {

		if dataPeserta[i].BidangMinat == "Cyber Security" || dataPeserta[i].BidangMinat == "Cyber" || dataPeserta[i].BidangMinat == "cyber" {
			jumlahCyber++
		}

		if dataPeserta[i].BidangMinat == "UI UX" || dataPeserta[i].BidangMinat == "UI" || dataPeserta[i].BidangMinat == "UX" || dataPeserta[i].BidangMinat == "ui" || dataPeserta[i].BidangMinat == "ux" {
			jumlahUIUX++
		}

		if dataPeserta[i].BidangMinat == "Data Science" {
			jumlahDataScience++
		}
	}
	fmt.Println(" ____________________ ")
	fmt.Println("|                    |")
	fmt.Println("| ################## |")
	fmt.Println("| Statistik KursusIn |")
	fmt.Println("| ################## |")
	fmt.Println("|____________________|")
	fmt.Println(" ____________________ ")
	fmt.Println("|Cyber Security :", jumlahCyber, " |")
	fmt.Println("|____________________| ")
	fmt.Println("|UI UX          :", jumlahUIUX, " |")
	fmt.Println("|____________________| ")
	fmt.Println("|Data Science   :", jumlahDataScience, " |")
	fmt.Println("|____________________| ")
	fmt.Println("|Total Peserta  :", len(dataPeserta), " |")
	fmt.Println("|____________________| ")
}

func tambahPeserta() {

	var p Peserta
	fmt.Print("masukan id : ")
	fmt.Scan(&p.ID)
	fmt.Print("masukan Nama : ")
	fmt.Scan(&p.Nama)
	fmt.Print("masukan Alamat : ")
	fmt.Scan(&p.Alamat)
	fmt.Print("masukan Minat : ")
	fmt.Scan(&p.BidangMinat)
	fmt.Print("masukan tgl masuk : ")
	fmt.Scan(&p.TanggalMasuk)

	dataPeserta = append(dataPeserta, p)
}

func hapusNama(nama string) {
	for i := 0; i < len(dataPeserta); i++ {

		if dataPeserta[i].Nama == nama {

			dataPeserta = append(
				dataPeserta[:i],
				dataPeserta[i+1:]...,
			)

			fmt.Println("Data berhasil dihapus")

			return
		}
	}

	fmt.Println("Nama tidak ditemukan")
}
func ubahPeserta1(nama string) {

	for i := 0; i < len(dataPeserta); i++ {

		if dataPeserta[i].Nama == nama {

			fmt.Print("Nama Baru: ")
			fmt.Scan(&dataPeserta[i].Nama)

			fmt.Print("Alamat Baru: ")
			fmt.Scan(&dataPeserta[i].Alamat)

			fmt.Println("Data berhasil diubah")

			return

		}
	}

	fmt.Println("Nama tidak ditemukan")
}
func ubahPeserta(id int) {

	for i := 0; i < len(dataPeserta); i++ {

		if dataPeserta[i].ID == id {

			fmt.Print("Nama Baru: ")
			fmt.Scan(&dataPeserta[i].Nama)

			fmt.Print("Alamat Baru: ")
			fmt.Scan(&dataPeserta[i].Alamat)

			fmt.Println("Data berhasil diubah")

			return

		}
	}

	fmt.Println("ID tidak ditemukan")
}
func sqnm(nama string) {
	for i := 0; i < len(dataPeserta); i++ {
		if dataPeserta[i].Nama == nama {
			fmt.Println("Data ditemukan")
			fmt.Println(dataPeserta[i])
			return
		}

	}
	fmt.Println("Data tidak ditemukan")
}

func sqid(id int) {
	for i := 0; i < len(dataPeserta); i++ {
		if dataPeserta[i].ID == id {
			fmt.Println("Data ditemukan")
			fmt.Println(dataPeserta[i])
			return
		}

	}
	fmt.Println("Data tidak ditemukan")
}

func hapusPeserta(id int) {

	for i := 0; i < len(dataPeserta); i++ {

		if dataPeserta[i].ID == id {

			dataPeserta = append(
				dataPeserta[:i],
				dataPeserta[i+1:]...,
			)

			fmt.Println("Data berhasil dihapus")

			return
		}
	}

	fmt.Println("ID tidak ditemukan")
}

func arr() {

	dataPeserta = append(dataPeserta,

		Peserta{
			ID:           1,
			Nama:         "Budi",
			Alamat:       "Cirebon",
			BidangMinat:  "Cyber Security",
			TanggalMasuk: "26-05-2026",
		},

		Peserta{
			ID:           2,
			Nama:         "Andi",
			Alamat:       "Bandung",
			BidangMinat:  "UI UX",
			TanggalMasuk: "27-05-2026",
		},

		Peserta{
			ID:           3,
			Nama:         "Rina",
			Alamat:       "Jakarta",
			BidangMinat:  "Data Science",
			TanggalMasuk: "28-05-2026",
		},

		Peserta{
			ID:           4,
			Nama:         "azzam",
			Alamat:       "Jakarta",
			BidangMinat:  "Data Science",
			TanggalMasuk: "28-07-2026",
		},

		Peserta{
			ID:           6,
			Nama:         "Rune",
			Alamat:       "Jakarta",
			BidangMinat:  "cyber",
			TanggalMasuk: "28-05-2026",
		},
		Peserta{
			ID:           5,
			Nama:         "Rinta",
			Alamat:       "Bandung",
			BidangMinat:  "Cyber",
			TanggalMasuk: "28-05-2026",
		},
	)

}

func tampil() {
	for i := 0; i < len(dataPeserta); i++ {

		fmt.Println(dataPeserta[i])
	}
}

func main() {

	arr()

	var pilih int

	for {
		fmt.Println(" ____________________ ")
		fmt.Println("|                    |")
		fmt.Println("| ################## |")
		fmt.Println("|      KursusIn      |")
		fmt.Println("|         by         |")
		fmt.Println("|   nopal & azka     |")
		fmt.Println("| ################## |")
		fmt.Println("|____________________|")

		fmt.Println("==========================")
		fmt.Println("Silahkan Pilih menu sesuai")
		fmt.Println("        keperluan         ")
		fmt.Println("==========================")

		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Menu Hapus dan ubah data")
		fmt.Println("3. Tampilkan Peserta")
		fmt.Println("4. mencari data sequential dan BINARY by ID OR NAME ")
		fmt.Println("5. statistik data dari kursus ini")
		fmt.Println("6. mengurutkan data by selectuion short & intersetion short ")
		fmt.Println("0. Keluar")

		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			tambahPeserta()

		case 2:

			var pilihHapus int

			fmt.Println("1. Hapus berdasarkan ID")
			fmt.Println("2. Hapus berdasarkan Nama")
			fmt.Println("3. ubah berdasarkan nama atau id")

			fmt.Println("4. Hapus semua data")

			fmt.Scan(&pilihHapus)

			switch pilihHapus {

			case 1:

				var id int

				fmt.Print("Masukkan ID: ")
				fmt.Scan(&id)

				hapusPeserta(id)

			case 2:

				var nama string

				fmt.Print("Masukkan Nama: ")
				fmt.Scan(&nama)

				hapusNama(nama)
			case 3:

				var input string

				fmt.Print("Masukkan Nama atau ID: ")
				fmt.Scan(&input)

				id, err := strconv.Atoi(input)

				if err == nil {

					ubahPeserta(id)

				} else {
					ubahPeserta1(input)
				}

			case 4:

				dataPeserta = nil

				fmt.Println("Semua data berhasil dihapus")

			default:
				fmt.Println("Pilihan tidak ada")
			}

		case 3:
			tampil()
		case 4:
			var pilih int

			fmt.Println("1. pilih sequential berdasarkan Nama dan ID")

			fmt.Println("2. pilih binary berdasarkan Nama dan ID")

			fmt.Scan(&pilih)

			switch pilih {
			case 1:
				var input string

				fmt.Print("Masukkan Nama atau ID: ")
				fmt.Scan(&input)

				id, err := strconv.Atoi(input)

				if err == nil {

					sqid(id)

				} else {
					sqnm(input)
				}
			case 2:

				var input string

				fmt.Print("Masukkan Nama atau ID: ")
				fmt.Scan(&input)

				id, err := strconv.Atoi(input)

				if err == nil {

					binarySearchID(id)

				} else {
					binarySearch(input)
				}
			}

		case 5:
			statistik()

		case 6:
			var pilih int

			fmt.Println("1. selection sort ")
			fmt.Println("2. insertion sort ")
			fmt.Println("3. selection plus binary ")
			fmt.Println("4. insertion plus binary ")

			fmt.Scan(&pilih)

			switch pilih {

			case 1:

				selectionSort()
				selectionSortNama()
				tampil()
				fmt.Println("data sudah di urutkan menggunakan selection")
			case 2:
				insertionSortNama()
				insertionSortId()
				tampil()
				fmt.Println("Data sudan diurutkan menggunakn intertion")

			case 3:
				var input string

				fmt.Print("Masukkan Nama atau ID: ")
				fmt.Scan(&input)

				id, err := strconv.Atoi(input)

				if err == nil {

					binarySearchID(id)
					selectionSort()

				} else {
					binarySearch(input)
					selectionSortNama()
				}
			case 4:
				var input string

				fmt.Print("Masukkan Nama atau ID: ")
				fmt.Scan(&input)

				id, err := strconv.Atoi(input)

				if err == nil {

					binarySearchID(id)
					insertionSortId()

				} else {
					binarySearch(input)
					insertionSortNama()
				}

			}

		case 0:
			return

		default:
			fmt.Println("Menu tidak ada")
		}
	}
}