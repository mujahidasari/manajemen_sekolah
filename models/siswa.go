package models

import "gorm.io/gorm"

type Siswa struct {
	gorm.Model
	Nama  string `json:"nama" binding:"required"`
	Kelas string `json:"kelas"`
	Email string `json:"email" gorm:"unique"`
}
