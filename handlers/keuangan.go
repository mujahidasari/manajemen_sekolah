package handlers

import (
	"manajemen-sekolah/db"
	"manajemen-sekolah/models"
	"manajemen-sekolah/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateKeuangan(c *gin.Context) {
	var keuangan models.Keuangan
	if err := c.ShouldBindJSON(&keuangan); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	// Simpan transaksi keuangan ke database
	if err := db.DB.Create(&keuangan).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create transaksi keuangan", err.Error())
		return
	}

	utils.SuccessResponse(c, "Transaksi keuangan created successfully", keuangan)
}

func GetAllKeuangan(c *gin.Context) {
	var keuangan []models.Keuangan

	// Filter berdasarkan siswa_id (opsional)
	siswaID := c.Query("siswa_id")
	if siswaID != "" {
		if err := db.DB.Where("siswa_id = ?", siswaID).Find(&keuangan).Error; err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
			return
		}
	} else {
		// Jika tidak ada filter, ambil semua data
		if err := db.DB.Find(&keuangan).Error; err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
			return
		}
	}

	utils.SuccessResponse(c, "Data fetched successfully", keuangan)
}

func UpdateKeuangan(c *gin.Context) {
	id := c.Param("id")
	var keuangan models.Keuangan

	// Cek apakah transaksi ada
	if err := db.DB.First(&keuangan, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Transaksi not found", err.Error())
		return
	}

	// Update transaksi berdasarkan input
	if err := c.ShouldBindJSON(&keuangan); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := db.DB.Save(&keuangan).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update transaksi", err.Error())
		return
	}

	utils.SuccessResponse(c, "Transaksi updated successfully", keuangan)
}

func DeleteKeuangan(c *gin.Context) {
	id := c.Param("id")
	var keuangan models.Keuangan

	// Cek apakah transaksi ada
	if err := db.DB.First(&keuangan, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Transaksi not found", err.Error())
		return
	}

	// Hapus transaksi
	if err := db.DB.Delete(&keuangan).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to delete Transaksi", err.Error())
		return
	}

	utils.SuccessResponse(c, "Transaksi deleted successfully", nil)
}
