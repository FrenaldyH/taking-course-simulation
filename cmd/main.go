package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/FrenaldyH/taking-course-simulation/config"
	"github.com/FrenaldyH/taking-course-simulation/internal/model"
	"github.com/FrenaldyH/taking-course-simulation/routes"
)

// @title			KRS API
// @version		1.0
// @description	API pengisian KRS sederhana untuk modul belajar backend.
// @description	Mahasiswa memilih mata kuliah dengan tiga aturan: tidak boleh ganda, kuota tidak boleh terlampaui, dan total SKS tidak boleh melebihi batas.
//
// @host			localhost:8080
// @BasePath		/
func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(
		&model.Mahasiswa{},
		&model.MataKuliah{},
		&model.KRS{},
	)
	if err != nil {
		log.Fatalf("Gagal menyiapkan tabel: %v", err)
	}

	config.SeedData()

	router := gin.Default()
	routes.RegisterRoutes(router)

	log.Println("Server jalan di http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server gagal jalan: %v", err)
	}
}
