package routes

import (
	"manajemen-sekolah/handlers"

	"github.com/gin-gonic/gin"
)

func SiswaRoutes(r *gin.Engine) {
	siswa := r.Group("/siswa")
	{
		siswa.GET("/", handlers.GetAllSiswa)
		siswa.POST("/", handlers.CreateSiswa)
		siswa.PUT("/:id", handlers.UpdateSiswa)
		siswa.DELETE("/:id", handlers.DeleteSiswa)
	}
}
