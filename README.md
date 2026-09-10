# KRS API

API pengisian Kartu Rencana Studi sederhana, dibangun dengan **Go**, **Gin**,
**GORM**, dan **PostgreSQL**.

Mahasiswa memilih mata kuliah, dan server memutuskan boleh atau tidak berdasarkan
tiga aturan:

1. Satu mata kuliah tidak boleh diambil dua kali
2. Kuota kelas tidak boleh terlampaui
3. Total SKS tidak boleh melebihi batas SKS mahasiswa

Materi belajar Go, Gin, GORM, PostgreSQL, Postman, dan Swagger yang membedah proyek
ini disusun sebagai modul terpisah.

<!-- TODO: pasang tautan ke modul di sini setelah alamatnya final. -->
<!-- Modul dibagikan sebagai ZIP dan dikelola pihak lain, jadi alamatnya belum ada. -->
<!-- Berkas yang perlu diubah kalau alamatnya sudah ada: baris ini dan bagian Struktur. -->


## Kebutuhan

- Go 1.25 atau lebih baru
- PostgreSQL yang sedang berjalan

## Menjalankan

```bash
cp .env.example .env    # Windows: copy .env.example .env
```

Isi kredensial PostgreSQL di berkas `.env`, lalu:

```bash
go mod download
go run ./cmd
```

Server berjalan di `http://localhost:8080`.

Saat pertama dijalankan, tabelnya dibuat otomatis lewat `AutoMigrate` dan diisi
data contoh:

```
Database connected successfully
Data contoh dimasukkan: 7 mata kuliah, 3 mahasiswa
Server jalan di http://localhost:8080
```

Dua data contoh dibuat agar aturan mudah dipicu: **IF2040** berkuota 2 kursi, dan
**Andi Wijaya** (id 3) berbatas 12 SKS.

## Endpoint


| Method   | Alamat               | Fungsi               |
| ---------- | ---------------------- | ---------------------- |
| `GET`    | `/ping`              | Cek server hidup     |
| `GET`    | `/mata-kuliah`       | Daftar mata kuliah   |
| `POST`   | `/krs`               | Ambil mata kuliah    |
| `GET`    | `/mahasiswa/:id/krs` | Lihat KRS mahasiswa  |
| `DELETE` | `/krs/:id`           | Batalkan mata kuliah |

Contoh mengambil mata kuliah:

```bash
curl -X POST http://localhost:8080/krs \
  -H 'Content-Type: application/json' \
  -d '{"mahasiswa_id": 1, "mata_kuliah_id": 1}'
```

Pelanggaran aturan dijawab `409 Conflict`, bukan `400`: permintaannya benar,
tapi keadaan saat itu yang menolak:

```json
{"error":"kuota mata kuliah sudah penuh (2 dari 2 kursi terisi)"}
```

## Struktur

```
.
├── cmd/
│   └── main.go                    Titik mulai program
├── routes/
│   └── routes.go                  Daftar seluruh alamat
├── internal/
│   ├── handler/                   Lapisan HTTP: baca request, kirim response
│   │   ├── krs_handler.go
│   │   ├── mata_kuliah_handler.go
│   │   └── response.go            Bentuk error + pemetaan status code
│   ├── service/                   Aturan bisnis
│   │   └── krs_service.go
│   ├── repo/                      Satu-satunya yang menyentuh database
│   │   ├── errors.go
│   │   ├── krs_repo.go
│   │   ├── mahasiswa_repo.go
│   │   └── mata_kuliah_repo.go
│   ├── model/                     Bentuk data
│   │   ├── krs.go
│   │   ├── mahasiswa.go
│   │   └── mata_kuliah.go
│   └── middleware/                Belum dipakai
├── config/
│   ├── database.go                Koneksi PostgreSQL
│   └── seed.go                    Data contoh saat pertama jalan
├── api-docs/                      Hasil generate Swagger, jangan diedit manual
├── postman/                       Koleksi request siap import
├── docs/schema/                   Skema database format DBML
└── pkg/                           Belum dipakai
```

Request mengalir satu arah dan tidak melompat:
`Router → Handler → Service → Repository → Database`.

Handler tidak menyentuh database langsung. Pemeriksaannya:

```bash
grep -rl "gorm.io" --include="*.go" .
```

Keluarannya hanya berkas di `config/` dan `internal/repo/`.

Pembahasan struktur ini berkas demi berkas ada di bab 08 modul.

<!-- TODO: tautkan ke bab 08 modul setelah alamatnya final. -->


## Swagger

Jalankan server, lalu buka http://localhost:8080/swagger/index.html

Dokumentasi dihasilkan dari komentar anotasi di atas tiap handler. Setiap kali
anotasi berubah, generate ulang:

```bash
swag init -g cmd/main.go -o ./api-docs --parseDependency --parseInternal
```

Tanda `-o ./api-docs` wajib. Tanpa itu, `swag` menulis ke `./docs` dan menimpa
folder skema.

CLI `swag` dipasang lebih dulu:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Postman

Import [`postman/krs-api.postman_collection.json`](postman/krs-api.postman_collection.json)
ke Postman: **Import** → seret berkasnya.

Isinya 13 request dalam 3 folder. Folder **Skenario Aturan** berisi urutan request
yang ditolak server, untuk menunjukkan ketiga aturan bisnis bekerja.

## Skema database

[`docs/schema/schema.dbml`](docs/schema/schema.dbml), dibuka dengan extension
**dbdiagram** di VS Code atau di https://dbdiagram.io

Tiga tabel: `mahasiswas`, `mata_kuliahs`, dan `krs` sebagai tabel penghubung.

Nama tabel berbentuk jamak karena dibuat GORM. Tabel `krs` dikunci namanya lewat
method `TableName()` di `internal/model/krs.go`.
