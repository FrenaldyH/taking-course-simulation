# Modul Backend: Go, Gin, dan GORM

Branch ini berisi contoh kode yang dipakai modul belajar backend.

**Materinya dibaca di wiki:**
https://github.com/FrenaldyH/taking-course-simulation/wiki

## Isi branch ini

| Folder | Dipakai di | Isi |
|---|---|---|
| `examples/01-hello` | Bab 02 | Variabel, tipe data, slice |
| `examples/02-struct` | Bab 02 | Struct, error handling, pointer |
| `examples/03-gin-ping` | Bab 03 | Server Gin paling minimal |
| `examples/04-gin-crud` | Bab 03, 04 | CRUD lengkap, data masih di variabel |
| `examples/05-gorm-connect` | Bab 06 | Koneksi PostgreSQL dan `AutoMigrate` |
| `examples/06-gorm-crud` | Bab 06 | CRUD, relasi, dan `Preload` |

## Menjalankan

```bash
go mod download
cp .env.example .env    # Windows: copy .env.example .env
go run ./examples/01-hello
```

Contoh 05 dan 06 memerlukan PostgreSQL yang sudah jalan dan kredensial yang benar
di berkas `.env`. Sisanya jalan tanpa persiapan tambahan.

## Proyek KRS

Kode yang dibedah di Bab 07 dan Bab 08 ada di branch `main` repo ini:
https://github.com/FrenaldyH/taking-course-simulation
