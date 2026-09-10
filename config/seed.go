package config

import (
	"log"

	"github.com/FrenaldyH/taking-course-simulation/internal/model"
)

// SeedData inserts sample rows on first run.
// It stops when data already exists, so restarts never duplicate rows.
func SeedData() {
	var jumlah int64
	DB.Model(&model.MataKuliah{}).Count(&jumlah)

	if jumlah > 0 {
		return
	}

	mataKuliah := []model.MataKuliah{
		{Kode: "IF2010", Nama: "Algoritma dan Pemrograman", SKS: 4, Kuota: 40},
		{Kode: "IF2020", Nama: "Struktur Data", SKS: 3, Kuota: 35},
		{Kode: "IF2030", Nama: "Basis Data", SKS: 3, Kuota: 30},
		{Kode: "IF2040", Nama: "Jaringan Komputer", SKS: 3, Kuota: 2},
		{Kode: "IF2050", Nama: "Sistem Operasi", SKS: 4, Kuota: 25},
		{Kode: "IF2060", Nama: "Rekayasa Perangkat Lunak", SKS: 3, Kuota: 30},
		{Kode: "IF2070", Nama: "Kecerdasan Buatan", SKS: 3, Kuota: 20},
	}

	mahasiswa := []model.Mahasiswa{
		{NIM: "5025231001", Nama: "Budi Santoso", BatasSKS: 24},
		{NIM: "5025231002", Nama: "Siti Rahayu", BatasSKS: 24},
		{NIM: "5025231003", Nama: "Andi Wijaya", BatasSKS: 12},
	}

	if err := DB.Create(&mataKuliah).Error; err != nil {
		log.Printf("Gagal mengisi data mata kuliah: %v", err)
		return
	}

	if err := DB.Create(&mahasiswa).Error; err != nil {
		log.Printf("Gagal mengisi data mahasiswa: %v", err)
		return
	}

	log.Printf("Data contoh dimasukkan: %d mata kuliah, %d mahasiswa",
		len(mataKuliah), len(mahasiswa))
}
