package repo

import (
	"errors"

	"gorm.io/gorm"

	"github.com/FrenaldyH/taking-course-simulation/config"
	"github.com/FrenaldyH/taking-course-simulation/internal/model"
)

// CreateKRS stores a new entry and fills its ID.
func CreateKRS(krs *model.KRS) error {
	return config.DB.Create(krs).Error
}

// FindKRSByMahasiswa returns a student's courses, with MataKuliah preloaded.
func FindKRSByMahasiswa(mahasiswaID uint) ([]model.KRS, error) {
	var daftar []model.KRS

	err := config.DB.
		Preload("MataKuliah").
		Where("mahasiswa_id = ?", mahasiswaID).
		Order("id").
		Find(&daftar).Error

	return daftar, err
}

// FindKRSByID returns one entry by id.
func FindKRSByID(id uint) (model.KRS, error) {
	var krs model.KRS

	err := config.DB.First(&krs, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.KRS{}, ErrNotFound
	}

	return krs, err
}

// DeleteKRS removes one entry.
// GORM does not error when nothing matched, so RowsAffected is checked here.
func DeleteKRS(id uint) error {
	hasil := config.DB.Delete(&model.KRS{}, id)
	if hasil.Error != nil {
		return hasil.Error
	}

	if hasil.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

// KRSExists reports whether the student already took this course.
func KRSExists(mahasiswaID, mataKuliahID uint) (bool, error) {
	var jumlah int64

	err := config.DB.Model(&model.KRS{}).
		Where("mahasiswa_id = ? AND mata_kuliah_id = ?", mahasiswaID, mataKuliahID).
		Count(&jumlah).Error

	return jumlah > 0, err
}

// TotalSKSDiambil sums a student's credits; COALESCE yields 0 when none.
func TotalSKSDiambil(mahasiswaID uint) (int, error) {
	var total int

	err := config.DB.Raw(`
		SELECT COALESCE(SUM(mk.sks), 0)
		FROM krs
		JOIN mata_kuliahs mk ON krs.mata_kuliah_id = mk.id
		WHERE krs.mahasiswa_id = ?`, mahasiswaID).Scan(&total).Error

	return total, err
}

// CountPeserta counts how many students already took one course.
func CountPeserta(mataKuliahID uint) (int, error) {
	var jumlah int64

	err := config.DB.Model(&model.KRS{}).
		Where("mata_kuliah_id = ?", mataKuliahID).
		Count(&jumlah).Error

	return int(jumlah), err
}
