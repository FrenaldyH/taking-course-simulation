package repo

import (
	"errors"

	"gorm.io/gorm"

	"github.com/FrenaldyH/taking-course-simulation/config"
	"github.com/FrenaldyH/taking-course-simulation/internal/model"
)

// FindAllMataKuliah returns every course, ordered by code.
func FindAllMataKuliah() ([]model.MataKuliah, error) {
	var daftar []model.MataKuliah
	err := config.DB.Order("kode").Find(&daftar).Error
	return daftar, err
}

// FindMataKuliahByID returns one course by id.
func FindMataKuliahByID(id uint) (model.MataKuliah, error) {
	var mk model.MataKuliah

	err := config.DB.First(&mk, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.MataKuliah{}, ErrNotFound
	}

	return mk, err
}
