package models

import "time"

//Item
type BarangKeluar struct {
	NoKwitansi   string
	SKU          *Barang
	Waktu        *time.Time
	JumlahKeluar int
	Harga        float64
	Catatan      string
}
