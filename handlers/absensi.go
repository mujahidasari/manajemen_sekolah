package handlers

import (
	"manajemen-sekolah/db"
	"manajemen-sekolah/models"
	"manajemen-sekolah/utils"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAbsensi(c *gin.Context) {
	var absensi models.Absensi
	if err := c.ShouldBindJSON(&absensi); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	// Validasi koordinat dengan lokasi sekolah (misalnya lat: -6.200000, lon: 106.816666)
	const schoolLat, schoolLon = -6.200000, 106.816666
	const radius = 500.0 // dalam meter
	if calculateDistance(absensi.Latitude, absensi.Longitude, schoolLat, schoolLon) > radius {
		utils.ErrorResponse(c, http.StatusBadRequest, "Location out of range", "")
		return
	}

	if err := db.DB.Create(&absensi).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create absensi", err.Error())
		return
	}

	utils.SuccessResponse(c, "Absensi berhasil dicatat", absensi)
}

func GetAllAbsensi(c *gin.Context) {
	var absensi []models.Absensi
	if err := db.DB.Find(&absensi).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
		return
	}
	utils.SuccessResponse(c, "Data fetched successfully", absensi)
}

// Menghitung jarak menggunakan Haversine Formula
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Radius bumi dalam meter
	phi1 := lat1 * (math.Pi / 180)
	phi2 := lat2 * (math.Pi / 180)
	deltaPhi := (lat2 - lat1) * (math.Pi / 180)
	deltaLambda := (lon2 - lon1) * (math.Pi / 180)

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c // Jarak dalam meter
}
