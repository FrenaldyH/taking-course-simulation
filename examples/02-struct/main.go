// Contoh 02: Struct, error handling, dan pointer.
// Jalankan: go run ./examples/02-struct
package main

import (
	"errors"
	"fmt"
)

// Mahasiswa is a blueprint for one student record.
// Every Mahasiswa value has exactly these fields.
type Mahasiswa struct {
	Nama     string
	NIM      string
	TotalSKS int
}

// cariMahasiswa returns a student, or an error when the NIM is unknown.
// Returning two values like this is the standard way Go reports failure.
func cariMahasiswa(nim string, daftar []Mahasiswa) (Mahasiswa, error) {
	for _, m := range daftar {
		if m.NIM == nim {
			return m, nil
		}
	}
	return Mahasiswa{}, errors.New("mahasiswa dengan NIM " + nim + " tidak ditemukan")
}

// tambahSKS receives a pointer, so it changes the original value
// instead of working on a copy.
func tambahSKS(m *Mahasiswa, jumlah int) {
	m.TotalSKS += jumlah
}

func main() {
	daftar := []Mahasiswa{
		{Nama: "Budi", NIM: "5025231001", TotalSKS: 20},
		{Nama: "Siti", NIM: "5025231002", TotalSKS: 22},
	}

	// The happy path: the NIM exists, so err is nil.
	budi, err := cariMahasiswa("5025231001", daftar)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Ketemu: %s (%s), %d SKS\n", budi.Nama, budi.NIM, budi.TotalSKS)

	// The failing path: the NIM does not exist, so err is filled.
	_, err = cariMahasiswa("9999999999", daftar)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Without a pointer, the function works on a copy and the original is untouched.
	fmt.Println("\nSebelum ditambah:", daftar[0].TotalSKS)
	tambahSKS(&daftar[0], 4)
	fmt.Println("Sesudah ditambah:", daftar[0].TotalSKS)
}
