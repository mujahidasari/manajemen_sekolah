package models

import "gorm.io/gorm"

type Tugas struct {
	gorm.Model
	GuruID     uint   `json:"guru_id"`
	Judul      string `json:"judul" binding:"required"`
	Deskripsi  string `json:"deskripsi"`
	Tanggal    string `json:"tanggal"`
	BatasWaktu string `json:"batas_waktu"`
}

type PengumpulanTugas struct {
	gorm.Model
	TugasID uint   `json:"tugas_id"`
	SiswaID uint   `json:"siswa_id"`
	File    string `json:"file"`
	Nilai   int    `json:"nilai"`
}
