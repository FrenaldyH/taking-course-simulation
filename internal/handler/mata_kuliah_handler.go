package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FrenaldyH/taking-course-simulation/internal/service"
)

// DaftarMataKuliah handles GET /mata-kuliah.
//
//	@Summary		Daftar mata kuliah
//	@Description	Menampilkan semua mata kuliah yang ditawarkan semester ini
//	@Tags			mata-kuliah
//	@Produce		json
//	@Success		200	{array}		internal_model.MataKuliah
//	@Failure		500	{object}	internal_handler.ErrorResponse
//	@Router			/mata-kuliah [get]
func DaftarMataKuliah(c *gin.Context) {
	daftar, err := service.DaftarMataKuliah()
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, daftar)
}
