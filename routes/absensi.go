package routes

import (
	"manajemen-sekolah/handlers"
	"manajemen-sekolah/middleware"

	"github.com/gin-gonic/gin"
)

func AbsensiRoutes(r *gin.Engine) {
	absensi := r.Group("/absensi")
	absensi.Use(middleware.JWTAuth())
	{
		absensi.POST("/", handlers.CreateAbsensi)
		absensi.GET("/", handlers.GetAllAbsensi)
	}
}
