package models

import "time"

//Item
type BarangKeluar struct {
	ID           int
	BarangID     int `form:"barang_id" binding:"required"`
	Barang       *Barang
	Waktu        time.Time
	JumlahKeluar int     `form:"jumlah_keluar" binding:"required"`
	Harga        float64 `form:"harga" binding:"required"`
	Catatan      string  `form:"catatan"`
}
