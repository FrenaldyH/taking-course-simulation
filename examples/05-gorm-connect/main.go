// Contoh 05: Menyambung ke PostgreSQL dan membuat tabel dari struct.
// Jalankan: go run ./examples/05-gorm-connect
// Butuh berkas .env berisi kredensial PostgreSQL.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MataKuliah becomes a table. GORM reads the struct tags to know
// which column rules to apply.
type MataKuliah struct {
	ID    uint   `gorm:"primaryKey"`
	Kode  string `gorm:"size:10;not null;unique"`
	Nama  string `gorm:"size:100;not null"`
	SKS   int    `gorm:"not null"`
	Kuota int    `gorm:"default:40"`
}

// getEnv returns an environment variable, or fallback when it is empty.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	// Read .env if it exists; system variables still work without it.
	if err := godotenv.Load(); err != nil {
		log.Println("Catatan: .env tidak ditemukan, memakai environment variable sistem")
	}

	// The DSN is the connection string PostgreSQL expects.
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
		log.Fatalf("Gagal menyambung ke database: %v\n"+
			"Pastikan PostgreSQL jalan dan isi .env sudah benar.", err)
	}

	fmt.Println("Berhasil tersambung ke database")

	// AutoMigrate creates the table when it does not exist,
	// and adds missing columns when the struct changes.
	if err := db.AutoMigrate(&MataKuliah{}); err != nil {
		log.Fatalf("Gagal membuat tabel: %v", err)
	}

	fmt.Println("Tabel 'mata_kuliahs' siap dipakai")

	// Ask the database itself which columns were actually created.
	kolom, err := db.Migrator().ColumnTypes(&MataKuliah{})
	if err != nil {
		log.Fatalf("Gagal membaca struktur tabel: %v", err)
	}

	fmt.Println("\nKolom yang dibuat GORM:")
	for _, k := range kolom {
		tipe, _ := k.ColumnType()
		fmt.Printf("  - %-8s %s\n", k.Name(), tipe)
	}
}
