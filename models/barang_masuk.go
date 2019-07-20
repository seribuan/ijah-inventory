package models

import "time"

//Item
type BarangMasuk struct {
	NoKwitansi     string
	SKU            *Barang
	Waktu          *time.Time
	JumlahPesan    int
	JumlahDiterima int
	Harga          float64
}
