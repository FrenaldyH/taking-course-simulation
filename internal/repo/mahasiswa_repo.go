package repo

import (
	"errors"

	"gorm.io/gorm"

	"github.com/FrenaldyH/taking-course-simulation/config"
	"github.com/FrenaldyH/taking-course-simulation/internal/model"
)

// FindMahasiswaByID returns one student by id.
func FindMahasiswaByID(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.First(&mahasiswa, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Mahasiswa{}, ErrNotFound
	}

	return mahasiswa, err
}

// FindAllMahasiswa returns every student, ordered by NIM.
func FindAllMahasiswa() ([]model.Mahasiswa, error) {
	var daftar []model.Mahasiswa
	err := config.DB.Order("nim").Find(&daftar).Error
	return daftar, err
}
