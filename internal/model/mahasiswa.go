package model

import "time"

// Mahasiswa is one student, with BatasSKS capping the credits they may take.
type Mahasiswa struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NIM       string    `json:"nim" gorm:"size:15;not null;unique"`
	Nama      string    `json:"nama" gorm:"size:100;not null"`
	BatasSKS  int       `json:"batas_sks" gorm:"not null;default:24"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
