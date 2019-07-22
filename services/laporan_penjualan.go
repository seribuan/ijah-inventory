package services

import (
	"sorabel/models"
	"time"

	"github.com/jinzhu/gorm"
)

type LaporanPenjualan struct {
	TanggalCetak   time.Time
	TotalOmzet     float64
	TotalLabaKotor float64
	TotalPenjualan float64
	TotalBarang    int
	Detail2        []PenjualanDetail
}

type PenjualanDetail struct {
	IDPesanan  string
	Waktu      time.Time
	SKU        string
	NamaBarang string
	Jumlah     int
	HargaJual  float64
	HargaBeli  float64
	Total      float64
	Laba       float64
}

func GenerateLaporanPenjualan(db *gorm.DB) LaporanPenjualan {
	var barang2Keluar []models.BarangKeluar
	var laporan LaporanPenjualan

	harga := GenerateHargaDanTotalNilaiBarang(db)

	db.Order("waktu").Preload("Barang").Find(&barang2Keluar)

	laporan.TanggalCetak = time.Now()
	for _, barangKeluar := range barang2Keluar {
		detail := generateDetailPenjualan(barangKeluar, harga[barangKeluar.BarangID])
		laporan.TotalOmzet += detail.Total
		laporan.TotalPenjualan++
		laporan.TotalBarang += detail.Jumlah
		laporan.TotalLabaKotor += detail.Laba
		laporan.Detail2 = append(laporan.Detail2, detail)
	}

	return laporan
}

func generateDetailPenjualan(barangKeluar models.BarangKeluar, harga HargaDanTotal) PenjualanDetail {
	var detail PenjualanDetail
	detail.IDPesanan = barangKeluar.Catatan
	detail.Waktu = barangKeluar.Waktu
	detail.SKU = barangKeluar.Barang.SKU
	detail.NamaBarang = barangKeluar.Barang.Nama
	detail.Jumlah = barangKeluar.JumlahKeluar
	detail.HargaJual = barangKeluar.Harga
	detail.Total = float64(barangKeluar.JumlahKeluar) * barangKeluar.Harga
	detail.HargaBeli = harga.Rata2Harga
	detail.Laba = detail.Total - (detail.HargaBeli * float64(detail.Jumlah))
	return detail
}
