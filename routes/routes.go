package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Blank import: registers the generated docs read by Swagger UI.
	_ "github.com/FrenaldyH/taking-course-simulation/api-docs"
	"github.com/FrenaldyH/taking-course-simulation/internal/handler"
)

// RegisterRoutes maps every URL to the handler that answers it.
func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/mata-kuliah", handler.DaftarMataKuliah)

	router.POST("/krs", handler.AmbilMataKuliah)
	router.DELETE("/krs/:id", handler.BatalkanKRS)

	router.GET("/mahasiswa/:id/krs", handler.LihatKRS)

	// Served at http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
