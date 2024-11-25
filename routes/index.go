package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	AuthRoutes(r)
	SiswaRoutes(r)
	AbsensiRoutes(r)
	TugasRoutes(r)
	KeuanganRoutes(r)
}
