# Catatan untuk Penerbit Modul

Berkas ini untuk yang mengatur distribusi modul, bukan untuk mahasiswa.
Hapus saja sebelum modul dibagikan.

## Yang perlu disesuaikan

Modul ini disiapkan tanpa mengetahui di mana ia akan diterbitkan, jadi ada dua
tempat yang memakai nama sementara `modul-backend-go`.

| Berkas | Baris | Isi sekarang | Perlu diapakan |
|---|---|---|---|
| `go.mod` | 3 | `module github.com/FrenaldyH/modul-backend-go` | Ganti kalau modul diterbitkan ke repo sendiri. Kalau tetap dibagikan sebagai ZIP, biarkan apa adanya |
| `00-persiapan.md` | 128, 135, 136 | `modul-backend-go` dan `<alamat-repo-modul>` | Sesuaikan dengan nama folder hasil ekstrak ZIP, atau isi alamat repo kalau diterbitkan |

Mengubah nama module di `go.mod` **aman**: keenam contoh di `examples/` tidak
saling impor, jadi tidak ada yang perlu ikut diubah.

`00-persiapan.md` bagian 5 sudah ditulis untuk dua kemungkinan sekaligus, ZIP
maupun repo. Kalau salah satu tidak dipakai, bagian itu boleh dipangkas.

## Yang sudah final, jangan diubah

Semua tautan ke repo proyek KRS sudah menunjuk alamat yang benar:

```
https://github.com/FrenaldyH/taking-course-simulation
```

Muncul di `README.md`, `07-swagger.md`, dan `08-studi-kasus-krs.md`. Repo itu
berisi kode yang dibedah di bab 08 dan dikelola terpisah dari modul ini.

## Memastikan modul masih utuh setelah diubah

```bash
go build ./...
go run ./examples/01-hello
```

Contoh 05 dan 06 memerlukan PostgreSQL dan berkas `.env`; sisanya jalan tanpa
persiapan tambahan.

## Isi modul

| Berkas | Keterangan |
|---|---|
| `README.md` | Indeks dan daftar isi |
| `00-persiapan.md` … `08-studi-kasus-krs.md` | Sembilan bab, dibaca berurutan |
| `CLI-CHEATSHEET.md` | Referensi perintah terminal |
| `examples/` | Enam contoh kode yang bisa dijalankan |
