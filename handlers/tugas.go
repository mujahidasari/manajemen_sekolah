package handlers

import (
	"manajemen-sekolah/db"
	"manajemen-sekolah/models"
	"manajemen-sekolah/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTugas(c *gin.Context) {
	var tugas models.Tugas
	if err := c.ShouldBindJSON(&tugas); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := db.DB.Create(&tugas).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create tugas", err.Error())
		return
	}

	utils.SuccessResponse(c, "Tugas created successfully", tugas)
}

func GetAllTugas(c *gin.Context) {
	var tugas []models.Tugas
	if err := db.DB.Find(&tugas).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
		return
	}
	utils.SuccessResponse(c, "Data fetched successfully", tugas)
}

func KumpulkanTugas(c *gin.Context) {
	// Logic untuk siswa mengumpulkan tugas
	c.JSON(http.StatusOK, gin.H{"message": "Tugas berhasil dikumpulkan"})
}
