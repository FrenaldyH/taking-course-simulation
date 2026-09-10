# Modul Backend dengan Go, Gin & GORM

Modul belajar backend untuk mahasiswa yang sudah paham logika pemrograman, tapi
belum pernah menyentuh backend, terminal, atau bahasa Go.

## Cara pakai

Baca berurutan dari bab 00. Tiap bab menumpuk di atas bab sebelumnya, dan istilah
baru dijelaskan saat pertama kali muncul.

Contoh kode ada di folder [`examples/`](examples/) dan bisa dijalankan langsung:

```bash
go run ./examples/01-hello
```

## Daftar isi

| Bab | Isi | Perkiraan baca |
|---|---|---|
| [00: Persiapan](00-persiapan.md) | Instalasi Go, PostgreSQL, VS Code, Postman | 45 menit, kerjakan sebelum kelas |
| [01: Backend Dasar](01-backend-dasar.md) | Client-server, HTTP, API, JSON | 20 menit |
| [02: Go Dasar](02-golang-dasar.md) | Sintaks Go seperlunya untuk backend | 30 menit |
| [03: Gin](03-gin.md) | Web server, routing, handler | 30 menit |
| [04: Postman](04-postman.md) | Menguji API: POST, PUT, DELETE | 15 menit |
| [05: PostgreSQL](05-postgresql.md) | Database relasional, tabel, SQL dasar | 25 menit |
| [06: GORM](06-gorm.md) | Menyambungkan Go ke database | 35 menit |
| [07: Swagger](07-swagger.md) | Dokumentasi API otomatis | 20 menit |
| [08: Studi Kasus KRS](08-studi-kasus-krs.md) | Bedah proyek nyata dan menambah fitur | 40 menit |

## Referensi

- [Cheat Sheet Terminal](CLI-CHEATSHEET.md), perintah yang dipakai sepanjang modul

Dua berkas berikut ada di repo proyek yang terpisah, karena keduanya menjelaskan
API yang dibedah di bab 08:

- [Koleksi Postman](https://github.com/FrenaldyH/taking-course-simulation/blob/main/postman/krs-api.postman_collection.json),
  13 request siap pakai termasuk skenario yang ditolak server
- [Skema database](https://github.com/FrenaldyH/taking-course-simulation/blob/main/docs/schema/schema.dbml),
  dibuka dengan extension dbdiagram di VS Code

## Alasan urutan bab

**Postman ditaruh setelah Gin.** Browser hanya bisa mengirim request GET. Begitu
endpoint yang menerima data (POST) dibuat, browser tidak lagi cukup.

**PostgreSQL ditaruh sebelum GORM.** GORM adalah alat untuk berbicara dengan
database. Tanpa mengetahui bentuk database yang diajak bicara, perintah seperti
`AutoMigrate` sulit dipahami dan sulit diperbaiki saat error.

## Proyek yang dibedah

Modul ini ditutup dengan API pengisian KRS: mahasiswa memilih mata kuliah, dengan
aturan batas SKS, kuota kelas, dan larangan mengambil mata kuliah yang sama dua
kali.

Kodenya ada di repo terpisah,
[taking-course-simulation](https://github.com/FrenaldyH/taking-course-simulation),
dan baru diperlukan mulai bab 07. Bab 00 sampai 06 cukup memakai contoh di repo ini.
