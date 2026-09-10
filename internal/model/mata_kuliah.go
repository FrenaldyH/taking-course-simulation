package model

import "time"

// MataKuliah is one course on offer, with Kuota capping how many may take it.
type MataKuliah struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Kode      string    `json:"kode" gorm:"size:10;not null;unique"`
	Nama      string    `json:"nama" gorm:"size:100;not null"`
	SKS       int       `json:"sks" gorm:"not null"`
	Kuota     int       `json:"kuota" gorm:"not null;default:40"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
