package service

import (
	"errors"
	"fmt"

	"github.com/FrenaldyH/taking-course-simulation/internal/model"
	"github.com/FrenaldyH/taking-course-simulation/internal/repo"
)

// Sentinel errors for the business rules; the handler maps them to status codes.
var (
	ErrMahasiswaTidakAda  = errors.New("mahasiswa tidak ditemukan")
	ErrMataKuliahTidakAda = errors.New("mata kuliah tidak ditemukan")
	ErrKRSTidakAda        = errors.New("data KRS tidak ditemukan")
	ErrSudahDiambil       = errors.New("mata kuliah ini sudah diambil")
	ErrKuotaPenuh         = errors.New("kuota mata kuliah sudah penuh")
	ErrMelebihiBatasSKS   = errors.New("melebihi batas SKS")
)

// DaftarMataKuliah returns every course on offer.
func DaftarMataKuliah() ([]model.MataKuliah, error) {
	return repo.FindAllMataKuliah()
}

// LihatKRS returns the courses a student has taken.
func LihatKRS(mahasiswaID uint) ([]model.KRS, error) {
	if _, err := repo.FindMahasiswaByID(mahasiswaID); err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return nil, ErrMahasiswaTidakAda
		}
		return nil, err
	}

	return repo.FindKRSByMahasiswa(mahasiswaID)
}

// AmbilMataKuliah registers a student into a course, rejecting duplicates,
// full classes, and totals above the student's credit limit.
func AmbilMataKuliah(mahasiswaID, mataKuliahID uint) (model.KRS, error) {
	mahasiswa, err := repo.FindMahasiswaByID(mahasiswaID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return model.KRS{}, ErrMahasiswaTidakAda
		}
		return model.KRS{}, err
	}

	mataKuliah, err := repo.FindMataKuliahByID(mataKuliahID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return model.KRS{}, ErrMataKuliahTidakAda
		}
		return model.KRS{}, err
	}

	// Checked first: taking the same course twice is the most common mistake.
	sudah, err := repo.KRSExists(mahasiswaID, mataKuliahID)
	if err != nil {
		return model.KRS{}, err
	}
	if sudah {
		return model.KRS{}, ErrSudahDiambil
	}

	peserta, err := repo.CountPeserta(mataKuliahID)
	if err != nil {
		return model.KRS{}, err
	}
	if peserta >= mataKuliah.Kuota {
		return model.KRS{}, fmt.Errorf("%w (%d dari %d kursi terisi)",
			ErrKuotaPenuh, peserta, mataKuliah.Kuota)
	}

	totalSKS, err := repo.TotalSKSDiambil(mahasiswaID)
	if err != nil {
		return model.KRS{}, err
	}
	if totalSKS+mataKuliah.SKS > mahasiswa.BatasSKS {
		return model.KRS{}, fmt.Errorf("%w: sudah %d SKS, menambah %d SKS, batas %d SKS",
			ErrMelebihiBatasSKS, totalSKS, mataKuliah.SKS, mahasiswa.BatasSKS)
	}

	krs := model.KRS{MahasiswaID: mahasiswaID, MataKuliahID: mataKuliahID}
	if err := repo.CreateKRS(&krs); err != nil {
		return model.KRS{}, err
	}

	// Attached so the caller gets the course details without a second request.
	krs.MataKuliah = &mataKuliah

	return krs, nil
}

// BatalkanKRS removes one course from a student's KRS.
func BatalkanKRS(id uint) error {
	err := repo.DeleteKRS(id)
	if errors.Is(err, repo.ErrNotFound) {
		return ErrKRSTidakAda
	}

	return err
}
