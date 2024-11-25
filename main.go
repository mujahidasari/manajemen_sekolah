package main

import (
	"manajemen-sekolah/db"
	"manajemen-sekolah/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.MigrateDB()

	r := gin.Default()

	// Registrasi semua routes
	routes.RegisterRoutes(r)

	r.Run(":8080") // Jalankan server pada port 8080
}
