[← Bab 01 Backend Dasar](01-backend-dasar.md) · [Daftar isi](README.md) · [Bab 03 Gin →](03-gin.md)

# Bab 02: Go Dasar

Bab ini **bukan** kursus Go lengkap. Isinya hanya bagian yang benar-benar dipakai
di bab 03 sampai 08.

Logika dasar seperti `if` dan perulangan sengaja dilewati. Fokusnya pada hal yang
**berbeda** dari bahasa lain: struct, cara Go menangani error, dan pointer.

## Struktur file Go paling minimal

```go
package main

import "fmt"

func main() {
	fmt.Println("Halo")
}
```

Tiga hal wajib:

- `package main`: menandakan berkas ini program yang bisa dijalankan
  (bukan pustaka yang dipakai program lain)
- `import`: daftar pustaka yang dipakai; `fmt` adalah pustaka bawaan untuk
  menulis ke layar
- `func main()`: titik mulai program. Selalu dari sini

Jalankan dengan:

```bash
go run nama_file.go
```

Go **menolak** meng-compile kalau ada `import` yang tidak dipakai. Extension Go
di VS Code membereskannya otomatis saat berkas disimpan.

## Variabel dan tipe statis

Jalankan contoh pertama:

```bash
go run ./examples/01-hello
```

Ada dua cara menulis variabel:

```go
var nama string = "Budi"   // lengkap: tipenya ditulis
nim := "5025231001"        // singkat: Go menebak tipenya sendiri
```

Bentuk singkat `:=` yang paling sering dipakai di dalam fungsi.

Go bertipe **statis**: tipe sebuah variabel ditetapkan sekali dan tidak bisa
berubah. Kalau `sks` sudah jadi angka, ia tidak bisa tiba-tiba diisi teks. Ini
berbeda dari Python atau JavaScript.

Konsekuensinya, Go tidak mau mencampur tipe diam-diam. Baris ini **gagal compile**:

```go
total := sks + nama   // error: mismatched types int and string
```

Konversinya harus ditulis sendiri:

```go
sksTeks := fmt.Sprintf("%d", sks)
```

Aturan ini membuat banyak bug ketahuan saat compile, bukan saat aplikasi sudah
dipakai.

### Zero value

Variabel yang dideklarasikan tanpa nilai **tidak berisi sampah acak**. Go mengisinya
dengan nilai nol sesuai tipenya:

| Tipe | Nilai awal |
|---|---|
| `int` | `0` |
| `string` | `""` (teks kosong) |
| `bool` | `false` |

### Slice: daftar yang bisa bertambah

```go
mataKuliah := []string{"Alpro", "Struktur Data"}
mataKuliah = append(mataKuliah, "Basis Data")

for i, mk := range mataKuliah {
	fmt.Printf("  %d. %s\n", i+1, mk)
}
```

Hasil `append` **harus ditampung kembali** ke variabelnya. Kalau tidak, datanya
seolah tidak bertambah.

## Struct: yang paling penting di modul ini

Struct adalah cetakan untuk data. Di bab 06, satu struct menjadi satu tabel di
database.

```go
type Mahasiswa struct {
	Nama     string
	NIM      string
	TotalSKS int
}
```

Membuat dan memakainya:

```go
budi := Mahasiswa{Nama: "Budi", NIM: "5025231001", TotalSKS: 20}
fmt.Println(budi.Nama)
```

Dalam istilah OOP, struct mirip class tapi hanya berisi data. Go tidak punya
pewarisan.

**Huruf besar di awal nama field itu bermakna, bukan sekadar gaya penulisan.**
Nama yang diawali huruf besar (`Nama`) bisa diakses dari package lain; yang diawali
huruf kecil (`nama`) hanya bisa diakses dari dalam package itu sendiri. Karena
nanti GORM dan Gin perlu membaca field ini dari package berbeda, field struct kita
**selalu** diawali huruf besar.

## Error handling ala Go

**Go tidak punya `try`/`catch`.**

Sebagai gantinya, fungsi yang bisa gagal mengembalikan **dua nilai**: hasilnya, dan
errornya.

```go
func cariMahasiswa(nim string, daftar []Mahasiswa) (Mahasiswa, error) {
	for _, m := range daftar {
		if m.NIM == nim {
			return m, nil          // berhasil: error diisi nil
		}
	}
	return Mahasiswa{}, errors.New("mahasiswa tidak ditemukan")   // gagal
}
```

Yang memanggil wajib memeriksanya:

```go
budi, err := cariMahasiswa("5025231001", daftar)
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println("Ketemu:", budi.Nama)
```

Pola `if err != nil` muncul **puluhan kali** sepanjang modul. Dibanding
`try`/`catch`, pola ini membuat setiap kemungkinan kegagalan terlihat di
tempatnya, bukan melompat ke penangan yang jauh.

`nil` artinya "kosong / tidak ada". `err != nil` berarti "ada error".

Jalankan contohnya:

```bash
go run ./examples/02-struct
```

Keluarannya:

```
Ketemu: Budi (5025231001), 20 SKS
Error: mahasiswa dengan NIM 9999999999 tidak ditemukan

Sebelum ditambah: 20
Sesudah ditambah: 24
```

## Pointer, seperlunya saja

Secara bawaan, saat sebuah nilai dikirim ke fungsi, Go mengirim **salinannya**.
Perubahan di dalam fungsi tidak berpengaruh ke aslinya.

Untuk mengubah nilai aslinya, kirim **alamatnya**. Itulah pointer:

```go
// Tanda * berarti "alamat menuju Mahasiswa", bukan Mahasiswa-nya langsung
func tambahSKS(m *Mahasiswa, jumlah int) {
	m.TotalSKS += jumlah
}

tambahSKS(&daftar[0], 4)   // tanda & berarti "ambil alamat dari"
```

Dua tanda:

| Tanda | Artinya |
|---|---|
| `&nilai` | Ambil alamat dari sebuah nilai |
| `*Tipe` | Tipe yang berisi alamat, bukan nilainya langsung |

Sebatas itu yang diperlukan modul ini. Di bab 06, `db.Create(&mahasiswa)` ditulis
dengan `&` karena GORM perlu **mengubah** struct itu untuk mengisi ID dari
database, jadi ia butuh alamat aslinya, bukan salinan.

## Mengelola pustaka

```bash
go mod init nama-project    # sekali saja, saat memulai project baru
go get nama-pustaka         # mengunduh pustaka dari internet
go mod tidy                 # merapikan: buang yang tak terpakai, tambah yang kurang
```

Hasilnya tercatat di dua berkas:

- `go.mod`: daftar pustaka yang dipakai beserta versinya
- `go.sum`: sidik jari tiap pustaka, untuk memastikan yang diunduh tidak berubah diam-diam

Keduanya **ikut di-commit** ke Git, supaya siapa pun yang meng-clone mendapat versi
pustaka yang sama persis.

## gofmt: merapikan kode otomatis

Go punya satu format resmi, dan alat untuk menerapkannya:

```bash
gofmt -w .        # rapikan semua berkas
gofmt -l .        # cuma daftar berkas yang belum rapi
```

Di VS Code dengan extension Go, ini berjalan otomatis setiap kali menyimpan.

Karena formatnya seragam, kode Go dari mana pun terlihat serupa dan tidak ada
perdebatan soal tab versus spasi.

---

## Ringkasan

| Hal | Intinya |
|---|---|
| `package main` + `func main()` | Titik mulai program |
| `:=` | Cara singkat membuat variabel |
| Tipe statis | Tipe tidak bisa berubah; konversi ditulis manual |
| Zero value | Variabel kosong berisi `0`, `""`, atau `false`: bukan sampah |
| **Struct** | Cetakan data; nanti jadi tabel di database |
| Huruf besar di awal | Menentukan bisa/tidaknya diakses package lain, bukan gaya penulisan |
| **`if err != nil`** | Cara Go menangani error; tidak ada try/catch |
| **Pointer `&`** | Dipakai kalau fungsi perlu mengubah nilai aslinya |
| `go get`, `go mod tidy` | Mengelola pustaka |

Yang dicetak tebal muncul terus sampai bab terakhir.

---

[← Bab 01 Backend Dasar](01-backend-dasar.md) · [Daftar isi](README.md) · [Bab 03 Gin →](03-gin.md)
