[← Bab 05 PostgreSQL](05-postgresql.md) · [Daftar isi](README.md) · [Bab 07 Swagger →](07-swagger.md)

# Bab 06: GORM

Bab 05 mengajarkan SQL. Bab ini mengajarkan cara menulis SQL **tanpa menulis SQL**.

## Apa itu ORM?

ORM (Object-Relational Mapping) adalah penerjemah antara struct Go dan tabel database.

Tanpa ORM, mengambil data berarti menulis SQL sebagai teks, lalu memindahkan
hasilnya ke struct satu kolom demi satu kolom. Dengan ORM:

```go
var daftar []MataKuliah
db.Find(&daftar)
```

Dua baris itu menjalankan `SELECT * FROM mata_kuliahs` dan sekaligus mengisi
slice-nya.

| Tanpa ORM | Dengan GORM |
|---|---|
| SQL ditulis sebagai teks, salah ketik baru ketahuan saat dijalankan | Ditulis sebagai kode Go, salah ketik ketahuan saat compile |
| Hasil query dipindah manual ke struct | Otomatis masuk ke struct |
| Ganti database berarti menulis ulang banyak query | Sebagian besar query tetap sama |

## Struct menjadi tabel

```go
type MataKuliah struct {
	ID    uint   `gorm:"primaryKey"`
	Kode  string `gorm:"size:10;not null;unique"`
	Nama  string `gorm:"size:100;not null"`
	SKS   int    `gorm:"not null"`
	Kuota int    `gorm:"default:40"`
}
```

Bandingkan dengan `CREATE TABLE` di bab 05. Isinya sama persis, hanya cara
menulisnya berbeda.

### Aturan penamaan yang mengejutkan

GORM mengubah nama secara otomatis:

| Di Go | Menjadi di database | Aturan |
|---|---|---|
| `MataKuliah` (nama struct) | tabel `mata_kuliahs` | huruf kecil, dipisah `_`, **ditambah `s`** |
| `Kode` (field) | kolom `kode` | huruf kecil |
| `BatasSKS` (field) | kolom `batas_sks` | dipisah `_` |

Nama tabelnya `mata_kuliahs`, bukan `mata_kuliah`, karena GORM memakai bentuk
jamak. Ini penyebab paling umum tabel tidak ketemu saat dicari di pgAdmin.

Kalau ingin menentukan nama sendiri:

```go
func (MataKuliah) TableName() string {
	return "mata_kuliah"
}
```

### Tag GORM yang sering dipakai

| Tag | Artinya | Padanan SQL |
|---|---|---|
| `primaryKey` | Kolom penanda unik | `PRIMARY KEY` |
| `not null` | Wajib diisi | `NOT NULL` |
| `unique` | Tidak boleh kembar | `UNIQUE` |
| `size:100` | Panjang maksimum teks | `VARCHAR(100)` |
| `default:40` | Nilai bawaan | `DEFAULT 40` |

Beberapa tag digabung dengan titik koma: `gorm:"size:10;not null;unique"`.

## Menyambung ke database

```bash
go run ./examples/05-gorm-connect
```

Inti kodenya:

```go
dsn := fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	host, user, password, dbname, port,
)

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
	log.Fatalf("Gagal menyambung ke database: %v", err)
}
```

**DSN** (Data Source Name) adalah satu baris teks berisi semua keterangan yang
dibutuhkan untuk menyambung: alamat, pengguna, sandi, nama database, port.

Kredensial dibaca dari berkas `.env`, bukan ditulis langsung di kode. `.env`
masuk `.gitignore`, sehingga password asli tidak ikut ter-upload ke GitHub.

## AutoMigrate

```go
db.AutoMigrate(&MataKuliah{})
```

Satu baris ini membaca struct, lalu menjalankan `CREATE TABLE` kalau tabelnya
belum ada.

Keluaran contoh 05 memperlihatkan kolom yang benar-benar terbentuk:

```
Berhasil tersambung ke database
Tabel 'mata_kuliahs' siap dipakai

Kolom yang dibuat GORM:
  - id       bigint
  - kode     character varying(10)
  - nama     character varying(100)
  - sks      bigint
  - kuota    bigint
```

AutoMigrate hanya bisa **menambah**:

| Perubahan di struct | AutoMigrate melakukannya? |
|---|---|
| Menambah field baru | Ya, kolom ditambahkan |
| Membuat tabel yang belum ada | Ya |
| Menghapus field | **Tidak**: kolomnya dibiarkan |
| Mengubah nama field | **Tidak**: dianggap kolom baru |

Sifat ini disengaja. Menghapus kolom berarti menghapus data, dan itu terlalu
berisiko untuk dilakukan otomatis.

## CRUD dengan GORM

```bash
go run ./examples/06-gorm-crud
```

### Create

```go
budi := Mahasiswa{NIM: "5025231001", Nama: "Budi", BatasSKS: 24}
db.Create(&budi)

fmt.Println(budi.ID)   // 1, diisi GORM setelah tersimpan
```

Di sinilah pointer dari bab 02 terpakai. Tanda `&` membuat GORM menerima
**alamat** struct, sehingga `ID` hasil dari database bisa dituliskan ke struct
aslinya. Tanpa `&`, GORM hanya mengisi salinan dan `budi.ID` tetap 0.

### Read

```go
var semua []MataKuliah
db.Find(&semua)                      // SELECT * FROM mata_kuliahs

var satu MataKuliah
db.First(&satu, 1)                   // ... WHERE id = 1 LIMIT 1

var tigaSKS []MataKuliah
db.Where("sks = ?", 3).Find(&tigaSKS)

db.Order("kode").Find(&semua)        // ORDER BY kode
```

**Tanda tanya `?` wajib dipakai**, bukan penggabungan teks:

```go
db.Where("sks = ?", nilai)                      // benar
db.Where("sks = " + nilai)                      // BAHAYA
```

Bentuk kedua membuka celah **SQL injection**, karena pengguna bisa menyisipkan
perintah SQL lewat isian biasa. Dengan `?`, GORM mengirim nilainya terpisah dari
perintah sehingga tidak ditafsirkan sebagai perintah.

### Data tidak ditemukan

```go
var mk MataKuliah
err := db.First(&mk, 9999).Error

if errors.Is(err, gorm.ErrRecordNotFound) {
	// tidak ada, biasanya dibalas 404
}
```

Saat data tidak ditemukan, GORM mencetak log berwarna merah:

```
/path/main.go:121 record not found
[0.456ms] [rows:0] SELECT * FROM "mata_kuliahs" WHERE id = 9999
```

Itu bukan crash, melainkan catatan dari logger GORM. Selama errornya diperiksa
dengan `errors.Is`, program berjalan normal.

### Update dan Delete

```go
satu.Kuota = 50
db.Save(&satu)                       // simpan seluruh isi struct

hasil := db.Delete(&MataKuliah{}, 1)
```

Ada satu perilaku Delete yang perlu diperhatikan:

```go
hasil := db.Delete(&KRS{}, 9999)     // id yang tidak ada
// hasil.Error        -> nil    (bukan error!)
// hasil.RowsAffected -> 0
```

Menghapus baris yang tidak ada **tidak dianggap error** oleh GORM. Untuk membalas
`404` pada kasus itu, `RowsAffected` harus diperiksa sendiri:

```go
if hasil.RowsAffected == 0 {
	return errors.New("data tidak ditemukan")
}
```

## Relasi antar tabel

Tabel penghubung `krs` dari bab 05 ditulis begini di GORM:

```go
type KRS struct {
	ID           uint `gorm:"primaryKey"`
	MahasiswaID  uint `gorm:"not null"`
	MataKuliahID uint `gorm:"not null"`

	Mahasiswa  Mahasiswa  `gorm:"foreignKey:MahasiswaID"`
	MataKuliah MataKuliah `gorm:"foreignKey:MataKuliahID"`
}
```

Dua kelompok field dengan tugas berbeda:

- `MahasiswaID`, `MataKuliahID`: **kolom asli** di database, berisi angka
- `Mahasiswa`, `MataKuliah`: **bukan kolom**; tempat GORM menaruh data terkait
  saat diminta

### Preload

Secara bawaan GORM **tidak** ikut mengambil data terkait:

```go
var k KRS
db.First(&k)
fmt.Println(k.MataKuliah.Nama)    // "" alias kosong!
```

Ini bukan bug. GORM tidak mengambil data terkait yang tidak diminta, agar
query-nya tetap ringan. Untuk memintanya:

```go
var krsBudi []KRS
db.Preload("MataKuliah").Where("mahasiswa_id = ?", 1).Find(&krsBudi)

for _, k := range krsBudi {
	fmt.Println(k.MataKuliah.Kode, k.MataKuliah.SKS)   // sekarang terisi
}
```

Keluaran contoh 06 memperlihatkan perbedaannya:

```
Tanpa Preload  -> nama mata kuliah: "" (kosong)
Dengan Preload -> KRS milik Budi:
  IF2010  Algoritma dan Pemrograman    4 SKS
  IF2020  Struktur Data                3 SKS
Total SKS diambil: 7 dari batas 24
```

Nama yang kosong padahal datanya ada di database adalah gejala khas `Preload`
yang terlupa.

## Kalau GORM tidak cukup

Untuk query rumit, SQL langsung sering lebih jelas:

```go
var total int
db.Raw(`
	SELECT COALESCE(SUM(mk.sks), 0)
	FROM krs
	JOIN mata_kuliahs mk ON krs.mata_kuliah_id = mk.id
	WHERE krs.mahasiswa_id = ?`, mahasiswaID).Scan(&total)
```

Di sinilah bab 05 terpakai: SQL bisa dibaca dan ditulis langsung saat ORM tidak
mencukupi.

---

## Ringkasan

| Hal | Intinya |
|---|---|
| Struct → tabel | Nama tabel jadi **jamak**: `MataKuliah` → `mata_kuliahs` |
| Tag `gorm:"..."` | Aturan kolom: `primaryKey`, `not null`, `unique`, `size`, `default` |
| `gorm.Open(...)` | Menyambung memakai DSN dari `.env` |
| `AutoMigrate` | Membuat tabel & menambah kolom, **tidak pernah menghapus** |
| `db.Create(&x)` | Simpan; `&` supaya ID bisa diisikan balik |
| `db.Find` / `First` / `Where` | Membaca data |
| `db.Where("sks = ?", n)` | **Selalu pakai `?`**: mencegah SQL injection |
| `gorm.ErrRecordNotFound` | Data tidak ada; log merahnya bukan crash |
| `RowsAffected` | Cara tahu Delete benar-benar menghapus sesuatu |
| `Preload("Nama")` | **Wajib** kalau ingin data relasinya ikut terisi |

---

[← Bab 05 PostgreSQL](05-postgresql.md) · [Daftar isi](README.md) · [Bab 07 Swagger →](07-swagger.md)
