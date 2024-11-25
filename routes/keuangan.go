package routes

import (
	"manajemen-sekolah/handlers"
	"manajemen-sekolah/middleware"

	"github.com/gin-gonic/gin"
)

func KeuanganRoutes(r *gin.Engine) {
	keuangan := r.Group("/keuangan")
	keuangan.Use(middleware.JWTAuth())
	{
		keuangan.POST("/", handlers.CreateKeuangan)      // Tambah transaksi
		keuangan.GET("/", handlers.GetAllKeuangan)       // List transaksi
		keuangan.PUT("/:id", handlers.UpdateKeuangan)    // Update transaksi
		keuangan.DELETE("/:id", handlers.DeleteKeuangan) // Hapus transaksi
	}
}
