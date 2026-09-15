// Contoh 01: Program Go pertama, variabel dan tipe data.
// Jalankan: go run ./examples/01-hello
package main

import "fmt"

func main() {
	// Full declaration: the type is written explicitly.
	var nama string = "Budi"

	// Short declaration: Go infers the type from the value.
	// This is the form used most often inside functions.
	nim := "5025231001"
	sks := 20
	lulus := true

	fmt.Println("Nama :", nama)
	fmt.Println("NIM  :", nim)
	fmt.Println("SKS  :", sks)
	fmt.Println("Lulus:", lulus)

	// A variable declared with a type but no value gets a "zero value",
	// never a random one: 0 for numbers, "" for strings, false for booleans.
	var belumDiisi int
	fmt.Println("Nilai awal int:", belumDiisi)

	// Go refuses to mix types silently. This would fail to compile:
	//   total := sks + nama
	// The conversion has to be written by hand.
	sksTeks := fmt.Sprintf("%d", sks)
	fmt.Println("SKS sebagai teks:", sksTeks+" SKS")

	// A slice is a list whose length can grow.
	mataKuliah := []string{"Alpro", "Struktur Data"}
	mataKuliah = append(mataKuliah, "Basis Data")

	fmt.Println("Jumlah mata kuliah:", len(mataKuliah))

	for i, mk := range mataKuliah {
		fmt.Printf("  %d. %s\n", i+1, mk)
	}
}
