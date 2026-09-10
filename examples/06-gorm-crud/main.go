// Contoh 06: CRUD dengan GORM, termasuk relasi antar tabel.
// Jalankan: go run ./examples/06-gorm-crud
// Butuh berkas .env berisi kredensial PostgreSQL.
package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID       uint   `gorm:"primaryKey"`
	NIM      string `gorm:"size:15;not null;unique"`
	Nama     string `gorm:"size:100;not null"`
	BatasSKS int    `gorm:"default:24"`
}

type MataKuliah struct {
	ID    uint   `gorm:"primaryKey"`
	Kode  string `gorm:"size:10;not null;unique"`
	Nama  string `gorm:"size:100;not null"`
	SKS   int    `gorm:"not null"`
	Kuota int    `gorm:"default:40"`
}

// KRS connects one Mahasiswa to one MataKuliah.
// The two extra fields let GORM load the related rows for us.
type KRS struct {
	ID           uint `gorm:"primaryKey"`
	MahasiswaID  uint `gorm:"not null"`
	MataKuliahID uint `gorm:"not null"`

	Mahasiswa  Mahasiswa  `gorm:"foreignKey:MahasiswaID"`
	MataKuliah MataKuliah `gorm:"foreignKey:MataKuliahID"`
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Catatan: .env tidak ditemukan, memakai environment variable sistem")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_NAME", "alpro_db"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal menyambung ke database: %v", err)
	}

	// Start from a clean state so this example prints the same thing every run.
	// Never do this in a real application.
	db.Migrator().DropTable(&KRS{}, &Mahasiswa{}, &MataKuliah{})

	if err := db.AutoMigrate(&Mahasiswa{}, &MataKuliah{}, &KRS{}); err != nil {
		log.Fatalf("Gagal membuat tabel: %v", err)
	}

	fmt.Println("== CREATE ==")

	budi := Mahasiswa{NIM: "5025231001", Nama: "Budi", BatasSKS: 24}

	// Pass the address so GORM can write the generated ID back into the struct.
	if err := db.Create(&budi).Error; err != nil {
		log.Fatalf("Gagal menyimpan mahasiswa: %v", err)
	}
	fmt.Printf("Mahasiswa tersimpan, ID diisi otomatis: %d\n", budi.ID)

	daftarMK := []MataKuliah{
		{Kode: "IF2010", Nama: "Algoritma dan Pemrograman", SKS: 4, Kuota: 40},
		{Kode: "IF2020", Nama: "Struktur Data", SKS: 3, Kuota: 35},
		{Kode: "IF2030", Nama: "Basis Data", SKS: 3, Kuota: 30},
	}
	if err := db.Create(&daftarMK).Error; err != nil {
		log.Fatalf("Gagal menyimpan mata kuliah: %v", err)
	}
	fmt.Printf("%d mata kuliah tersimpan\n", len(daftarMK))

	fmt.Println("\n== READ ==")

	var semua []MataKuliah
	db.Order("kode").Find(&semua)
	for _, mk := range semua {
		fmt.Printf("  %s  %-28s %d SKS\n", mk.Kode, mk.Nama, mk.SKS)
	}

	var satu MataKuliah
	if err := db.First(&satu, daftarMK[0].ID).Error; err != nil {
		log.Fatalf("Gagal membaca satu baris: %v", err)
	}
	fmt.Printf("First(id=%d): %s\n", daftarMK[0].ID, satu.Nama)

	var tigaSKS []MataKuliah
	db.Where("sks = ?", 3).Find(&tigaSKS)
	fmt.Printf("Mata kuliah 3 SKS: %d buah\n", len(tigaSKS))

	// Looking for something that does not exist returns a specific error.
	var hilang MataKuliah
	err = db.First(&hilang, 9999).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("First(id=9999): ErrRecordNotFound (memang tidak ada)")
	}

	fmt.Println("\n== UPDATE ==")

	satu.Kuota = 50
	db.Save(&satu)
	fmt.Printf("Kuota %s diubah jadi %d\n", satu.Kode, satu.Kuota)

	fmt.Println("\n== RELASI ==")

	db.Create(&KRS{MahasiswaID: budi.ID, MataKuliahID: daftarMK[0].ID})
	db.Create(&KRS{MahasiswaID: budi.ID, MataKuliahID: daftarMK[1].ID})

	// Without Preload, the Mahasiswa and MataKuliah fields stay empty.
	var tanpaPreload KRS
	db.First(&tanpaPreload)
	fmt.Printf("Tanpa Preload  -> nama mata kuliah: %q (kosong)\n", tanpaPreload.MataKuliah.Nama)

	// Preload tells GORM to fetch the related rows too.
	var krsBudi []KRS
	db.Preload("MataKuliah").Where("mahasiswa_id = ?", budi.ID).Find(&krsBudi)

	totalSKS := 0
	fmt.Println("Dengan Preload -> KRS milik Budi:")
	for _, k := range krsBudi {
		fmt.Printf("  %s  %-28s %d SKS\n", k.MataKuliah.Kode, k.MataKuliah.Nama, k.MataKuliah.SKS)
		totalSKS += k.MataKuliah.SKS
	}
	fmt.Printf("Total SKS diambil: %d dari batas %d\n", totalSKS, budi.BatasSKS)

	fmt.Println("\n== DELETE ==")

	hasil := db.Delete(&KRS{}, krsBudi[0].ID)
	fmt.Printf("Baris terhapus: %d\n", hasil.RowsAffected)

	// Delete succeeds even when nothing matched, so check RowsAffected yourself.
	hasil = db.Delete(&KRS{}, 9999)
	fmt.Printf("Hapus id tidak ada -> error: %v, RowsAffected: %d\n",
		hasil.Error, hasil.RowsAffected)
}
