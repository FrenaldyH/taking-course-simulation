package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FrenaldyH/taking-course-simulation/internal/service"
)

// ErrorResponse is the single error shape returned by every endpoint.
type ErrorResponse struct {
	Error string `json:"error" example:"mata kuliah tidak ditemukan"`
}

// MessageResponse confirms an action that returns no data.
type MessageResponse struct {
	Message string `json:"message" example:"berhasil dibatalkan"`
}

// parseID reads a numeric URL parameter, writing the 400 response itself
// and reporting whether parsing succeeded.
func parseID(c *gin.Context, nama string) (uint, bool) {
	nilai, err := strconv.ParseUint(c.Param(nama), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: nama + " harus berupa angka"})
		return 0, false
	}

	return uint(nilai), true
}

// respondError maps a service error to its HTTP status code.
func respondError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrMahasiswaTidakAda),
		errors.Is(err, service.ErrMataKuliahTidakAda),
		errors.Is(err, service.ErrKRSTidakAda):
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})

	// Conflict, not BadRequest: the request is well formed but the current
	// state rejects it.
	case errors.Is(err, service.ErrSudahDiambil),
		errors.Is(err, service.ErrKuotaPenuh),
		errors.Is(err, service.ErrMelebihiBatasSKS):
		c.JSON(http.StatusConflict, ErrorResponse{Error: err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}
}
