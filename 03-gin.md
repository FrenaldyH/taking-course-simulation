[← Bab 02 Go Dasar](02-golang-dasar.md) · [Daftar isi](README.md) · [Bab 04 Postman →](04-postman.md)

# Bab 03: Gin

Bab ini membangun web server pertama yang bisa menjawab request dari browser.

## Apa itu Gin

Go punya pustaka bawaan untuk membuat web server, `net/http`. Dengan pustaka itu
saja, hal yang selalu dibutuhkan harus ditulis berulang: mencocokkan URL, membaca
angka dari alamat, mengubah struct jadi JSON. Gin menyediakannya sudah jadi.

| Tanpa Gin (`net/http`) | Dengan Gin |
|---|---|
| Cocokkan URL manual dengan `if` dan pemotongan teks | `router.GET("/mata-kuliah/:id", ...)` |
| Ubah data ke JSON manual | `c.JSON(200, data)` |
| Baca body request manual | `c.ShouldBindJSON(&data)` |

Gin bukan satu-satunya pilihan; ada juga Echo, Fiber, dan Chi. Gin dipakai di
modul ini karena paling banyak digunakan, sehingga contoh di internet paling
melimpah.

## Server pertama

```bash
go run ./examples/03-gin-ping
```

Isinya:

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	log.Println("Server jalan di http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server gagal jalan: %v", err)
	}
}
```

Buka http://localhost:8080/ping di browser. Hasilnya:

```json
{"message":"pong"}
```

Empat baris intinya:

| Baris | Artinya |
|---|---|
| `gin.Default()` | Membuat "mesin" yang menerima semua request |
| `router.GET("/ping", ...)` | Mendaftarkan satu alamat beserta cara menjawabnya |
| `c.JSON(...)` | Mengirim balasan dalam bentuk JSON |
| `router.Run(":8080")` | Menyalakan server di port 8080 |

`router.Run` **menahan program di baris itu** selamanya. Terminal tidak kembali
ke prompt, sesuai penjelasan di bab 01. Hentikan dengan `Ctrl+C`.

`gin.H` adalah cara singkat menulis objek JSON. `gin.H{"message": "pong"}`
menghasilkan `{"message":"pong"}`.

## Handler dan gin.Context

Fungsi yang menjawab sebuah alamat disebut **handler**. Bentuknya selalu sama:

```go
func namaHandler(c *gin.Context) {
	// ...
}
```

Satu-satunya bekal handler adalah `c`, yaitu `gin.Context`. Di dalamnya ada
request yang masuk, dan lewat objek itu pula response dikirim balik.

| Untuk membaca request | |
|---|---|
| `c.Param("id")` | Ambil bagian dari URL, misal `/mata-kuliah/3` |
| `c.Query("kode")` | Ambil dari query string, misal `?kode=IF2010` |
| `c.ShouldBindJSON(&x)` | Baca body JSON, masukkan ke struct `x` |

| Untuk mengirim response | |
|---|---|
| `c.JSON(status, data)` | Kirim JSON beserta status code-nya |

### Mengambil nilai dari URL

Titik dua menandai bagian yang berubah-ubah:

```go
router.GET("/halo/:nama", func(c *gin.Context) {
	nama := c.Param("nama")
	c.JSON(http.StatusOK, gin.H{"pesan": "Halo, " + nama})
})
```

Buka http://localhost:8080/halo/Budi:

```json
{"pesan":"Halo, Budi"}
```

## Status code jangan asal 200

Gin menyediakan nama yang mudah dibaca, jadi tidak perlu menghafal angkanya:

| Konstanta | Angka | Dipakai saat |
|---|---|---|
| `http.StatusOK` | 200 | Berhasil |
| `http.StatusCreated` | 201 | Data baru berhasil dibuat |
| `http.StatusBadRequest` | 400 | Kiriman client tidak valid |
| `http.StatusNotFound` | 404 | Data yang dicari tidak ada |
| `http.StatusInternalServerError` | 500 | Kode kita yang bermasalah |

Menulis `http.StatusNotFound` lebih baik daripada `404` karena maksudnya langsung
terbaca tanpa mengingat arti angkanya.

## CRUD lengkap

```bash
go run ./examples/04-gin-crud
```

Contoh ini punya lima alamat:

| Method | Alamat | Fungsi |
|---|---|---|
| `GET` | `/mata-kuliah` | Lihat semua |
| `GET` | `/mata-kuliah/:id` | Lihat satu |
| `POST` | `/mata-kuliah` | Tambah baru |
| `PUT` | `/mata-kuliah/:id` | Ubah |
| `DELETE` | `/mata-kuliah/:id` | Hapus |

### Membaca data yang dikirim client

```go
func tambah(c *gin.Context) {
	var baru MataKuliah

	if err := c.ShouldBindJSON(&baru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON tidak valid"})
		return
	}
	// ...
}
```

Tiga hal yang terjadi:

1. `var baru MataKuliah` menyiapkan wadah kosong
2. `&baru` mengirim **alamatnya** (pointer, bab 02). Gin perlu **mengisi** struct
   itu, jadi butuh alamat aslinya
3. Kalau JSON kiriman rusak, `err` terisi, dan kita balas `400` lalu `return`

`return` di situ **wajib**. Tanpa itu, kode di bawahnya tetap jalan dan server
mengirim dua response sekaligus.

### Tag json menentukan nama field

```go
type MataKuliah struct {
	ID   int    `json:"id"`
	Kode string `json:"kode"`
	Nama string `json:"nama"`
	SKS  int    `json:"sks"`
}
```

Tanpa tag `json:"..."`, Go akan mengeluarkan nama field apa adanya, yaitu `SKS`, bukan
`sks`. Tag inilah yang menjembatani gaya penamaan Go (huruf besar di awal) dengan
gaya penamaan JSON (huruf kecil).

## Mencoba sendiri

`GET` bisa langsung dari browser:

```
http://localhost:8080/mata-kuliah
```

Hasilnya:

```json
[{"id":1,"kode":"IF2010","nama":"Algoritma dan Pemrograman","sks":4},
 {"id":2,"kode":"IF2020","nama":"Struktur Data","sks":3}]
```

Untuk `POST`, `PUT`, dan `DELETE`, **browser tidak cukup**. Mengetik alamat di
address bar selalu menghasilkan request `GET`, dan tidak ada cara mengirim body
JSON dari sana. Alat penggantinya dibahas di [bab 04](04-postman.md).

Hasil sebenarnya dari kelima alamat:

| Request | Response | Status |
|---|---|---|
| `POST /mata-kuliah` body `{"kode":"IF2030","nama":"Basis Data","sks":3}` | `{"id":3,"kode":"IF2030","nama":"Basis Data","sks":3}` | 201 |
| `GET /mata-kuliah/3` | `{"id":3,"kode":"IF2030","nama":"Basis Data","sks":3}` | 200 |
| `PUT /mata-kuliah/3` body `{"kode":"IF2030","nama":"Basis Data Lanjut","sks":4}` | `{"id":3,"kode":"IF2030","nama":"Basis Data Lanjut","sks":4}` | 200 |
| `DELETE /mata-kuliah/3` | `{"message":"mata kuliah dihapus"}` | 200 |
| `GET /mata-kuliah/abc` | `{"error":"id harus berupa angka"}` | 400 |
| `GET /mata-kuliah/999` | `{"error":"mata kuliah tidak ditemukan"}` | 404 |

## Masalah yang sengaja dibiarkan

Langkahnya:

1. Jalankan `go run ./examples/04-gin-crud`
2. Tambah satu mata kuliah lewat `POST`
3. Tekan `Ctrl+C` untuk mematikan server
4. Jalankan lagi, lalu buka `GET /mata-kuliah`

Mata kuliah yang baru ditambahkan **hilang**. Yang tersisa hanya dua data awal
yang ditulis di dalam kode.

Penyebabnya ada di baris ini:

```go
var daftar = []MataKuliah{ ... }
```

Data hidup di dalam variabel, dan variabel hidup di memori (RAM). Begitu program
berhenti, memorinya dilepas.

Ada masalah kedua yang lebih halus: server melayani banyak request **bersamaan**.
Kalau dua orang menambah data pada saat yang sama, satu slice biasa seperti ini
bisa rusak datanya.

Dua masalah itu, data tidak awet dan tidak aman diakses bersamaan, diselesaikan
oleh database. Dibahas di [bab 05](05-postgresql.md).

---

## Ringkasan

| Hal | Intinya |
|---|---|
| `gin.Default()` | Membuat mesin penerima request |
| `router.GET/POST/PUT/DELETE` | Mendaftarkan alamat |
| `c *gin.Context` | Satu-satunya bekal handler: baca request, kirim response |
| `c.Param("id")` | Ambil bagian berubah-ubah dari URL |
| `c.ShouldBindJSON(&x)` | Baca body JSON ke dalam struct |
| `c.JSON(status, data)` | Kirim balasan |
| Tag `json:"nama"` | Menentukan nama field di JSON |
| `return` setelah kirim error | Wajib, supaya tidak mengirim dua response |
| Data di variabel | Hilang saat server mati, alasan kita butuh database |

---

[← Bab 02 Go Dasar](02-golang-dasar.md) · [Daftar isi](README.md) · [Bab 04 Postman →](04-postman.md)
