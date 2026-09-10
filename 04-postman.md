[← Bab 03 Gin](03-gin.md) · [Daftar isi](README.md) · [Bab 05 PostgreSQL →](05-postgresql.md)

# Bab 04: Postman

## Kenapa perlu alat terpisah?

Di bab 03, `GET /mata-kuliah` bisa dibuka lewat browser. Untuk `POST`, browser
tidak cukup.

Mengetik alamat di address bar **selalu** menghasilkan request `GET`. Tidak ada
cara mengirim `POST` beserta body JSON dari sana.

Padahal tiga dari empat operasi CRUD butuh itu:

| Operasi | Bisa dari browser? |
|---|---|
| `GET`: lihat data | Bisa |
| `POST`: tambah data | Tidak |
| `PUT`: ubah data | Tidak |
| `DELETE`: hapus data | Tidak |

Postman bisa mengirim keempatnya, lengkap dengan body dan header yang ditentukan
sendiri.

## Siapkan dulu

Jalankan server contoh dari bab 03, biarkan hidup di terminal terpisah:

```bash
go run ./examples/04-gin-crud
```

Semua percobaan di bab ini mengarah ke server tersebut.

## Bagian-bagian jendela Postman

Saat membuat request baru, ada empat bagian:

| Bagian | Letaknya | Gunanya |
|---|---|---|
| **Method** | Dropdown di kiri kolom alamat | Memilih `GET` / `POST` / `PUT` / `DELETE` |
| **URL** | Kolom panjang di tengah | Alamat tujuan |
| **Body** | Tab di bawah kolom alamat | Data yang dikirim (hanya untuk POST/PUT) |
| **Response** | Panel bawah | Jawaban server + status code |

## 1. Request pertama: GET

1. Klik **New** → **HTTP Request** (atau tanda `+` di tab)
2. Biarkan method-nya `GET`
3. Isi URL: `http://localhost:8080/mata-kuliah`
4. Klik **Send**

Di panel bawah akan muncul:

```json
[
    {
        "id": 1,
        "kode": "IF2010",
        "nama": "Algoritma dan Pemrograman",
        "sks": 4
    },
    {
        "id": 2,
        "kode": "IF2020",
        "nama": "Struktur Data",
        "sks": 3
    }
]
```

Perhatikan juga baris kecil di kanan atas panel response:

```
Status: 200 OK    Time: 5 ms    Size: 231 B
```

Angka `200` adalah status code dari bab 01. Bagian ini yang pertama diperiksa
saat request tidak berjalan sesuai harapan.

## 2. Mengirim data baru: POST

Operasi ini tidak bisa dilakukan browser.

1. Buat request baru
2. Ubah method menjadi **POST**
3. URL: `http://localhost:8080/mata-kuliah`
4. Buka tab **Body**
5. Pilih **raw**
6. Di dropdown sebelah kanannya, ganti `Text` menjadi **JSON**
7. Ketik isinya:

```json
{
    "kode": "IF2030",
    "nama": "Basis Data",
    "sks": 3
}
```

8. Klik **Send**

Hasilnya:

```json
{
    "id": 3,
    "kode": "IF2030",
    "nama": "Basis Data",
    "sks": 3
}
```

Statusnya **201 Created**, bukan 200: ada data baru yang berhasil dibuat, dan
`id`-nya diisi otomatis oleh server.

**Langkah 6 paling sering terlewat.** Kalau dropdown masih `Text`, Postman tidak
menambahkan header `Content-Type: application/json`, dan server menolak dengan
`400 JSON tidak valid`. Periksa bagian ini lebih dulu saat POST gagal padahal
JSON-nya terlihat benar.

## 3. Mengubah data: PUT

1. Method: **PUT**
2. URL: `http://localhost:8080/mata-kuliah/3`: perhatikan ada `/3` di belakang
3. Body → raw → JSON:

```json
{
    "kode": "IF2030",
    "nama": "Basis Data Lanjut",
    "sks": 4
}
```

Hasilnya:

```json
{
    "id": 3,
    "kode": "IF2030",
    "nama": "Basis Data Lanjut",
    "sks": 4
}
```

Bedanya dengan POST, alamatnya menyebut **data yang mana** yang diubah (`/3`).

## 4. Menghapus data: DELETE

1. Method: **DELETE**
2. URL: `http://localhost:8080/mata-kuliah/3`
3. Body tidak perlu diisi
4. Send

```json
{
    "message": "mata kuliah dihapus"
}
```

Kirim ulang `GET /mata-kuliah` untuk memastikan: data `id: 3` sudah tidak ada.

## 5. Melihat error dengan sengaja

Tiga request berikut sengaja ditolak server:

| Request | Response | Status |
|---|---|---|
| `GET /mata-kuliah/abc` | `{"error":"id harus berupa angka"}` | `400` |
| `GET /mata-kuliah/999` | `{"error":"mata kuliah tidak ditemukan"}` | `404` |
| `POST /mata-kuliah` dengan body `{"kode":"","nama":""}` | `{"error":"kode dan nama wajib diisi"}` | `400` |

Ketiganya `4xx`, artinya **kiriman client yang bermasalah**, bukan server yang
rusak. Status `500` berarti kesalahan ada di kode Go.

## 6. Menyimpan request ke Collection

Agar URL tidak perlu diketik ulang setiap kali:

1. Klik **Save** pada sebuah request (atau `Ctrl+S`)
2. Klik **New Collection**, beri nama misalnya `Belajar Backend`
3. Simpan

Request berikutnya disimpan ke koleksi yang sama, lalu bisa dikirim ulang kapan
saja tanpa mengetik apa pun.

### Variabel supaya tidak mengetik alamat berulang

Kalau alamat server berubah, mengubahnya satu per satu di puluhan request akan
merepotkan. Pakai variabel:

1. Klik nama collection → tab **Variables**
2. Tambah variabel `base_url` dengan nilai `http://localhost:8080`
3. Di request, tulis alamatnya jadi `{{base_url}}/mata-kuliah`

Dengan begitu, satu perubahan berlaku untuk semua request.

## Kesalahan yang paling sering terjadi

| Gejala | Penyebab | Perbaikan |
|---|---|---|
| `Could not send request` / `ECONNREFUSED` | Servernya belum jalan | Jalankan `go run ...` di terminal, biarkan terbuka |
| `400 JSON tidak valid` padahal JSON terlihat benar | Body masih `Text`, belum `JSON` | Ganti dropdown di tab Body |
| `400 JSON tidak valid` | Ada koma berlebih setelah field terakhir | Hapus koma terakhir |
| `404 page not found` | Alamat salah ketik, atau method-nya keliru | Cocokkan dengan daftar alamat di kode |
| Data hilang setelah server di-restart | Memang belum pakai database | Bukan bug, dijelaskan di bab 05 |

Ada dua jenis 404. `{"error":"mata kuliah tidak ditemukan"}` berarti alamatnya
benar tapi datanya tidak ada. `404 page not found` berarti **alamatnya** yang
tidak dikenali server.

---

## Ringkasan

- Browser hanya bisa mengirim `GET`; Postman bisa keempat method
- Untuk POST/PUT: tab **Body** → **raw** → ganti `Text` jadi **JSON**
- Selalu baca **status code** lebih dulu saat menelusuri masalah
- `4xx` = kiriman client salah, `5xx` = kode server bermasalah
- Simpan request ke **Collection** supaya tidak mengetik ulang
- Pakai variabel `{{base_url}}` supaya alamat mudah diganti serentak

---

[← Bab 03 Gin](03-gin.md) · [Daftar isi](README.md) · [Bab 05 PostgreSQL →](05-postgresql.md)
