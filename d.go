package main

import "fmt"

const NMAX int = 100

type peserta struct {
	id   int
	nama string
}

type arrPeserta [NMAX]peserta

func tampilData(T arrPeserta, n int) {
	var i int

	i = 0
	for i < n {
		fmt.Println(T[i].id, T[i].nama)
		i = i + 1
	}
}

func selectionSortID(T *arrPeserta, n int) {
	
	var i, j, idx_min int
	var t peserta

	i = 1
	for i <= n-1 {

		idx_min = i - 1
		j = i

		for j < n {

			if T[idx_min].id > T[j].id {
				idx_min = j
			}

			j = j + 1
		}

		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t

		i = i + 1
	}
}

func insertionSortNama(T *arrPeserta, n int) {
	
	var i, j int
	var x peserta

	i = 1

	for i <= n-1 {

		j = i
		x = T[j]

		for j > 0 && x.nama > T[j-1].nama {

			T[j] = T[j-1]

			j = j - 1
		}

		T[j] = x

		i = i + 1
	}
}

func main() {

	var data arrPeserta
	var n int

	n = 4

	data[0].id = 103
	data[0].nama = "Rina"

	data[1].id = 101
	data[1].nama = "Budi"

	data[2].id = 105
	data[2].nama = "Angga"

	data[3].id = 102
	data[3].nama = "Citra"

	fmt.Println("Data Awal")
	tampilData(data, n)

	fmt.Println()
	fmt.Println("Selection Sort ID Ascending")
	selectionSortID(&data, n)
	tampilData(data, n)

	fmt.Println()
	fmt.Println("Insertion Sort Nama Descending")
	insertionSortNama(&data, n)
	tampilData(data, n)
}