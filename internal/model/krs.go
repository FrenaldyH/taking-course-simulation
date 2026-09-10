package model

import "time"

// KRS records that one student takes one course.
// Mahasiswa and MataKuliah are not columns: GORM fills them on Preload,
// and omitempty keeps them out of the JSON when they were not loaded.
type KRS struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	MahasiswaID  uint `json:"mahasiswa_id" gorm:"not null;index"`
	MataKuliahID uint `json:"mata_kuliah_id" gorm:"not null;index"`

	Mahasiswa  *Mahasiswa  `json:"mahasiswa,omitempty" gorm:"foreignKey:MahasiswaID"`
	MataKuliah *MataKuliah `json:"mata_kuliah,omitempty" gorm:"foreignKey:MataKuliahID"`

	CreatedAt time.Time `json:"created_at"`
}

// TableName pins the table to "krs"; GORM would otherwise pluralize the
// acronym and break the raw SQL in the repository layer.
func (KRS) TableName() string {
	return "krs"
}
