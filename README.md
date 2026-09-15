# Materi 3 - Contoh Kode Backend

Branch `materi` berisi contoh kode per bab untuk materi backend
(Go, Gin, GORM, PostgreSQL).

**Materi lengkapnya di wiki:**
[Materi 3 ‐ Backend](https://github.com/Algoritma-dan-Pemrograman-ITS/LBE-2026/wiki/Materi-3-%E2%80%90-Backend)

## Isi branch ini

| Folder | Dipakai di | Isi |
|---|---|---|
| `examples/01-hello` | Bab 02 | Variabel, tipe data, slice |
| `examples/02-struct` | Bab 02 | Struct, error handling, pointer |
| `examples/03-gin-ping` | Bab 03 | Server Gin paling minimal |
| `examples/04-gin-crud` | Bab 03, 04 | CRUD lengkap, data masih di variabel |
| `examples/05-gorm-connect` | Bab 06 | Koneksi PostgreSQL dan `AutoMigrate` |
| `examples/06-gorm-crud` | Bab 06 | CRUD, relasi, dan `Preload` |

## Mengunduh

```bash
git clone -b materi --single-branch https://github.com/Algoritma-dan-Pemrograman-ITS/LBE-2026.git lbe-materi
cd lbe-materi
```

## Menjalankan

```bash
go mod download
cp .env.example .env    # Windows: copy .env.example .env
go run ./examples/01-hello
```

Contoh 05 dan 06 memerlukan PostgreSQL yang sudah jalan dan kredensial yang benar
di berkas `.env`. Sisanya jalan tanpa persiapan tambahan.

## Proyek KRS

Kode yang dibedah di Bab 07 dan Bab 08 ada di branch
[`project-backend`](https://github.com/Algoritma-dan-Pemrograman-ITS/LBE-2026/tree/project-backend).
