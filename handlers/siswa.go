package handlers

import (
	"manajemen-sekolah/db"
	"manajemen-sekolah/models"
	"manajemen-sekolah/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSiswa(c *gin.Context) {
	var siswa []models.Siswa
	if err := db.DB.Find(&siswa).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
		return
	}
	utils.SuccessResponse(c, "Data fetched successfully", siswa)
}

func CreateSiswa(c *gin.Context) {
	var siswa models.Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := db.DB.Create(&siswa).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create siswa", err.Error())
		return
	}

	utils.SuccessResponse(c, "Siswa created successfully", siswa)
}

func UpdateSiswa(c *gin.Context) {
	id := c.Param("id")
	var siswa models.Siswa

	if err := db.DB.First(&siswa, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Siswa not found", err.Error())
		return
	}

	if err := c.ShouldBindJSON(&siswa); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := db.DB.Save(&siswa).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update siswa", err.Error())
		return
	}

	utils.SuccessResponse(c, "Siswa updated successfully", siswa)
}

func DeleteSiswa(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Siswa{}, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete siswa", err.Error())
		return
	}

	utils.SuccessResponse(c, "Siswa deleted successfully", nil)
}
