# Cheat Sheet Terminal: Perintah yang Dipakai Hari Ini

Simpan halaman ini terbuka selama sesi. Semua perintah di sini akan kita pakai.

---

## Cara membuka terminal

| OS | Cara |
|---|---|
| Windows | Tekan `Win`, ketik "PowerShell", Enter |
| macOS | Tekan `Cmd+Space`, ketik "Terminal", Enter |
| Linux | Tekan `Ctrl+Alt+T` |
| Di dalam VSCode (paling praktis) | Tekan `` Ctrl+` `` (tombol backtick, di bawah Esc) |

---

## Navigasi folder

| Perintah | Arti | Contoh |
|---|---|---|
| `pwd` | Sekarang saya ada di folder mana? | `pwd` |
| `ls` | Lihat isi folder ini | `ls` |
| `cd nama_folder` | Masuk ke folder | `cd modul-backend-go` |
| `cd ..` | Naik satu tingkat ke folder induk | `cd ..` |

> Windows PowerShell: `ls` dan `pwd` juga jalan. Kalau pakai Command Prompt lama,
> gunakan `dir` (ganti `ls`) dan `cd` (ganti `pwd`).

Ketik 2-3 huruf awal nama folder lalu tekan **Tab**, terminal akan melengkapi
sendiri. Cara ini menghindari salah ketik.

---

## Perintah Go

| Perintah | Arti |
|---|---|
| `go version` | Cek Go sudah terinstall dan versinya berapa |
| `go run cmd/main.go` | Jalankan program (yang paling sering kita pakai hari ini) |
| `go get nama-package` | Download library dari internet |
| `go mod tidy` | Rapikan daftar library: buang yang tak terpakai, tambah yang kurang |
| `go build ./...` | Cek semua kode bisa di-compile (tanpa menjalankannya) |

---

## Mengontrol program yang sedang jalan

| Tombol | Fungsi |
|---|---|
| `Ctrl+C` | **Hentikan server.** Server tidak berhenti sendiri |
| `↑` (panah atas) | Panggil ulang perintah sebelumnya, tak perlu ketik ulang |

> Kalau muncul error **"port 8080 already in use"**, artinya masih ada server lama
> yang jalan. Cari terminal yang menjalankannya, tekan `Ctrl+C`.

---

## Perintah Git

| Perintah | Arti |
|---|---|
| `git clone <url>` | Unduh project dari GitHub |
| `git status` | Lihat berkas apa saja yang berubah |
| `git pull` | Ambil pembaruan terbaru dari GitHub |

Kode versi lengkap ada di repo, sehingga bisa dibandingkan setelah sesi selesai.

---

## Membaca pesan error

Pesan error di terminal adalah petunjuk. Yang perlu dicari:

1. **Baris paling atas**: biasanya penyebab sebenarnya
2. **Nama file + nomor baris**, contoh `main.go:15` → buka file itu, lihat baris 15

Error yang sering muncul:

| Pesan error | Artinya |
|---|---|
| `undefined: xxx` | Menulis nama fungsi/variabel yang belum dibuat, atau salah ketik |
| `expected ';'` / `syntax error` | Ada kurung `{` `}` yang belum ditutup |
| `connection refused` | PostgreSQL belum jalan |
| `password authentication failed` | Password di file `.env` salah |
| `404 page not found` | **Bukan error**: server hidup, cuma route-nya belum dibuat |
