package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FrenaldyH/taking-course-simulation/internal/service"
)

// AmbilMataKuliahInput is the request body for taking a course.
type AmbilMataKuliahInput struct {
	MahasiswaID  uint `json:"mahasiswa_id" example:"1"`
	MataKuliahID uint `json:"mata_kuliah_id" example:"2"`
}

// AmbilMataKuliah handles POST /krs.
//
//	@Summary		Ambil mata kuliah
//	@Description	Mendaftarkan mahasiswa ke sebuah mata kuliah. Ditolak jika mata kuliah sudah diambil, kuota penuh, atau melebihi batas SKS.
//	@Tags			krs
//	@Accept			json
//	@Produce		json
//	@Param			input	body		AmbilMataKuliahInput	true	"Mahasiswa dan mata kuliah yang dipilih"
//	@Success		201		{object}	internal_model.KRS
//	@Failure		400		{object}	internal_handler.ErrorResponse
//	@Failure		404		{object}	internal_handler.ErrorResponse
//	@Failure		409		{object}	internal_handler.ErrorResponse
//	@Router			/krs [post]
func AmbilMataKuliah(c *gin.Context) {
	var input AmbilMataKuliahInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "JSON tidak valid"})
		return
	}

	if input.MahasiswaID == 0 || input.MataKuliahID == 0 {
		c.JSON(http.StatusBadRequest,
			ErrorResponse{Error: "mahasiswa_id dan mata_kuliah_id wajib diisi"})
		return
	}

	krs, err := service.AmbilMataKuliah(input.MahasiswaID, input.MataKuliahID)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, krs)
}

// LihatKRS handles GET /mahasiswa/:id/krs.
//
//	@Summary		Lihat KRS mahasiswa
//	@Description	Menampilkan semua mata kuliah yang sudah diambil seorang mahasiswa
//	@Tags			krs
//	@Produce		json
//	@Param			id	path		int	true	"ID mahasiswa"
//	@Success		200	{array}		internal_model.KRS
//	@Failure		400	{object}	internal_handler.ErrorResponse
//	@Failure		404	{object}	internal_handler.ErrorResponse
//	@Router			/mahasiswa/{id}/krs [get]
func LihatKRS(c *gin.Context) {
	mahasiswaID, ok := parseID(c, "id")
	if !ok {
		return
	}

	daftar, err := service.LihatKRS(mahasiswaID)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, daftar)
}

// BatalkanKRS handles DELETE /krs/:id.
//
//	@Summary		Batalkan mata kuliah
//	@Description	Menghapus satu mata kuliah dari KRS mahasiswa
//	@Tags			krs
//	@Produce		json
//	@Param			id	path		int	true	"ID baris KRS"
//	@Success		200	{object}	internal_handler.MessageResponse
//	@Failure		400	{object}	internal_handler.ErrorResponse
//	@Failure		404	{object}	internal_handler.ErrorResponse
//	@Router			/krs/{id} [delete]
func BatalkanKRS(c *gin.Context) {
	krsID, ok := parseID(c, "id")
	if !ok {
		return
	}

	if err := service.BatalkanKRS(krsID); err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: "mata kuliah berhasil dibatalkan"})
}
