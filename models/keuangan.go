package models

import "gorm.io/gorm"

type Keuangan struct {
	gorm.Model
	SiswaID    uint   `json:"siswa_id"`             // Referensi ke siswa
	Jumlah     int    `json:"jumlah"`               // Jumlah transaksi
	Status     string `json:"status"`               // Status pembayaran (Lunas/Belum Lunas)
	Keterangan string `json:"keterangan,omitempty"` // Opsional: Deskripsi transaksi
}
