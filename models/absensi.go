package models

import "gorm.io/gorm"

type Absensi struct {
	gorm.Model
	SiswaID   uint    `json:"siswa_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Foto      string  `json:"foto"`
	Status    string  `json:"status"` // Hadir, Izin, Sakit
}
