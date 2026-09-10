// Contoh 04: CRUD lengkap, tapi data disimpan di variabel (belum ada database).
// Jalankan: go run ./examples/04-gin-crud
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MataKuliah is the shape of one course.
// The json tags decide the field names in the JSON output.
type MataKuliah struct {
	ID   int    `json:"id"`
	Kode string `json:"kode"`
	Nama string `json:"nama"`
	SKS  int    `json:"sks"`
}

// Data lives in memory only, so it disappears when the program stops.
// A real server also handles requests in parallel, which makes a plain
// slice like this unsafe. One more reason to use a database instead.
var (
	daftar = []MataKuliah{
		{ID: 1, Kode: "IF2010", Nama: "Algoritma dan Pemrograman", SKS: 4},
		{ID: 2, Kode: "IF2020", Nama: "Struktur Data", SKS: 3},
	}
	idBerikutnya = 3
)

func main() {
	router := gin.Default()

	router.GET("/mata-kuliah", ambilSemua)
	router.GET("/mata-kuliah/:id", ambilSatu)
	router.POST("/mata-kuliah", tambah)
	router.PUT("/mata-kuliah/:id", ubah)
	router.DELETE("/mata-kuliah/:id", hapus)

	log.Println("Server jalan di http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server gagal jalan: %v", err)
	}
}

// ambilSemua handles GET /mata-kuliah
func ambilSemua(c *gin.Context) {
	c.JSON(http.StatusOK, daftar)
}

// ambilSatu handles GET /mata-kuliah/:id
func ambilSatu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus berupa angka"})
		return
	}

	for _, mk := range daftar {
		if mk.ID == id {
			c.JSON(http.StatusOK, mk)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "mata kuliah tidak ditemukan"})
}

// tambah handles POST /mata-kuliah
func tambah(c *gin.Context) {
	var baru MataKuliah

	// ShouldBindJSON reads the request body and fills the struct.
	if err := c.ShouldBindJSON(&baru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON tidak valid"})
		return
	}

	if baru.Kode == "" || baru.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kode dan nama wajib diisi"})
		return
	}

	baru.ID = idBerikutnya
	idBerikutnya++
	daftar = append(daftar, baru)

	c.JSON(http.StatusCreated, baru)
}

// ubah handles PUT /mata-kuliah/:id
func ubah(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus berupa angka"})
		return
	}

	var ubahan MataKuliah
	if err := c.ShouldBindJSON(&ubahan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON tidak valid"})
		return
	}

	// Loop by index so the change is written back into the slice.
	for i := range daftar {
		if daftar[i].ID == id {
			daftar[i].Kode = ubahan.Kode
			daftar[i].Nama = ubahan.Nama
			daftar[i].SKS = ubahan.SKS

			c.JSON(http.StatusOK, daftar[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "mata kuliah tidak ditemukan"})
}

// hapus handles DELETE /mata-kuliah/:id
func hapus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id harus berupa angka"})
		return
	}

	for i, mk := range daftar {
		if mk.ID == id {
			// Remove index i by joining the parts before and after it.
			daftar = append(daftar[:i], daftar[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"message": "mata kuliah dihapus"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "mata kuliah tidak ditemukan"})
}
