package models

import "time"

//Item
type BarangMasuk struct {
	ID             int
	BarangID       int `form:"barang_id" binding:"required"`
	Barang         *Barang
	Waktu          time.Time
	JumlahPesan    int     `form:"jumlah_pesan" binding:"required"`
	JumlahDiterima int     `form:"jumlah_diterima"`
	Harga          float64 `form:"harga" binding:"required"`
	NoKwitansi     string  `form:"no_kwitansi" binding:"required"`
	Catatan        string  `form:"catatan"`
}
