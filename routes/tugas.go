package routes

import (
	"manajemen-sekolah/handlers"
	"manajemen-sekolah/middleware"

	"github.com/gin-gonic/gin"
)

func TugasRoutes(r *gin.Engine) {
	tugas := r.Group("/tugas")
	tugas.Use(middleware.JWTAuth())
	{
		tugas.POST("/", handlers.CreateTugas)
		tugas.GET("/", handlers.GetAllTugas)
		tugas.POST("/:id/kumpul", handlers.KumpulkanTugas)
	}
}
