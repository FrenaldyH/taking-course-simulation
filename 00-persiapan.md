[← Kembali ke daftar isi](README.md) · Selanjutnya: [01: Backend Dasar](01-backend-dasar.md)

# Bab 00: Persiapan


| Alat           | Gunanya                                 |
| ---------------- | ----------------------------------------- |
| **Go**         | Bahasa yang kita pakai menulis backend  |
| **VS Code**    | Tempat menulis kode                     |
| **PostgreSQL** | Database, tempat data disimpan permanen |
| **Postman**    | Alat untuk menguji API yang kita buat   |

---

## 1. Install Go

**Windows**

1. Unduh installer dari https://go.dev/dl/ (pilih berkas `.msi`)
2. Jalankan installer, klik Next sampai selesai
3. **Tutup semua terminal / PowerShell yang sedang terbuka**, lalu buka yang baru

**macOS**

1. Unduh berkas `.pkg` dari https://go.dev/dl/
2. Jalankan installer sampai selesai

**Linux (Ubuntu/Debian)**

```bash
sudo apt update && sudo apt install golang-go
```

### Verifikasi

Buka terminal, ketik:

```bash
go version
```

---

## 2. Install VS Code + Extension Go

1. Unduh VS Code: https://code.visualstudio.com/
2. Buka VS Code, tekan `Ctrl+Shift+X` (Mac: `Cmd+Shift+X`)
3. Cari **"Go"**, pasang extension resmi dari **Go Team at Google**
4. Buat berkas baru bernama `coba.go`, ketik `package main`
5. Kalau muncul notifikasi meminta memasang tools tambahan, klik **Install All**

Extension ini memberi tahu kesalahan ketik sebelum kode dijalankan.

### Verifikasi

Berkas `.go` yang dibuka punya pewarnaan sintaks: kata `package` berwarna berbeda
dari kata `main`.

---

## 3. Install PostgreSQL

**Windows / macOS**

1. Unduh dari https://www.postgresql.org/download/
2. Jalankan installer
3. **CATAT PASSWORD yang kamu buat.** Jangan sampai lupa
4. Port biarkan bawaan: `5432`
5. Sertakan **pgAdmin** saat instalasi (biasanya sudah tercentang)

**Linux (Ubuntu/Debian)**

```bash
sudo apt update && sudo apt install postgresql
sudo systemctl start postgresql
sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"
```

### Buat satu database kosong

Lewat pgAdmin: login dengan password tadi → klik kanan `Databases` → `Create` →
`Database` → beri nama **`alpro_db`** → Save.

Lewat terminal:

```bash
psql -U postgres -c "CREATE DATABASE alpro_db;"
```

### Verifikasi

Database `alpro_db` muncul di daftar sebelah kiri pgAdmin.

### Kalau error


| Pesan                            | Artinya                             | Perbaikan                                                                                            |
| ---------------------------------- | ------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| `connection refused`             | Layanan PostgreSQL belum jalan      | Windows: buka`services.msc`, cari `postgresql`, klik Start. Linux: `sudo systemctl start postgresql` |
| `password authentication failed` | Password salah                      | Ulangi instalasi atau reset password                                                                 |
| `port 5432 already in use`       | Sudah ada PostgreSQL lain terpasang | Pakai yang sudah ada, tidak perlu pasang dua                                                         |

---

## 4. Install Postman

Unduh di https://www.postman.com/downloads/

Postman dipakai untuk mengirim request ke API buatan kita. Browser hanya bisa
mengirim request jenis GET, sedangkan nanti kita perlu mengirim data baru (POST),
mengubah (PUT), dan menghapus (DELETE). Penjelasan lengkapnya ada di
[bab 04](04-postman.md).

Saat pertama dibuka, Postman menawarkan pembuatan akun. Untuk keperluan modul ini
akun **tidak diperlukan**: cari tautan kecil bertuliskan "Skip and go to the app"
atau semacamnya.

---

## 5. Ambil contoh-contoh kode

Seluruh contoh di modul ini ada dalam satu paket, dibagikan sebagai berkas ZIP
atau alamat repo.

Kalau menerima **berkas ZIP**, ekstrak lalu masuk ke foldernya:

```bash
cd modul-backend-go
go mod download
```

Kalau menerima **alamat repo**:

```bash
git clone <alamat-repo-modul>
cd modul-backend-go
go mod download
```

Lalu buat berkas `.env` dari templatnya:

```bash
cp .env.example .env
```

> Windows PowerShell memakai: `copy .env.example .env`

Buka berkas `.env` di VS Code, ganti `your_password_here` dengan password
PostgreSQL kamu.

### Verifikasi akhir

```bash
go run ./examples/01-hello
```

Harus muncul:

```
Nama : Budi
NIM  : 5025231001
SKS  : 20
Lulus: true
```

diikuti beberapa baris lagi.

Keluaran itu menandakan instalasi sudah benar: Go terpasang, pustaka terunduh, dan
program Go bisa dijalankan dari terminal.

**[Bab 01: Backend Dasar](01-backend-dasar.md)**.

---

[← Daftar isi](README.md) · [Cheat sheet terminal](CLI-CHEATSHEET.md) · [Bab 01 →](01-backend-dasar.md)
