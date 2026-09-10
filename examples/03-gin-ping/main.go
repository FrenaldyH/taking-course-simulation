// Contoh 03: Server Gin paling minimal.
// Jalankan: go run ./examples/03-gin-ping
// Lalu buka http://localhost:8080/ping di browser.
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.Default() creates the engine that receives every request.
	router := gin.Default()

	// Register one address: GET /ping
	router.GET("/ping", func(c *gin.Context) {
		// gin.H is a shortcut for building a JSON object.
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Read a value from the URL, for example /halo/Budi
	router.GET("/halo/:nama", func(c *gin.Context) {
		nama := c.Param("nama")
		c.JSON(http.StatusOK, gin.H{"pesan": "Halo, " + nama})
	})

	log.Println("Server jalan di http://localhost:8080")

	// Run blocks here and keeps waiting for requests until Ctrl+C.
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server gagal jalan: %v", err)
	}
}
