package services

import (
	"sorabel/models"
	"time"

	"github.com/jinzhu/gorm"
)

type LaporanPenjualan struct {
	TanggalCetak    time.Time
	JumlahSKU       int
	TotalOmzet      int
	TotalLabaKotor  float64
	TotalPenjualan  float64
	TotalBarang     int
	PenjualanDetail PenjualanDetail
}

type PenjualanDetail struct {
	IDBarang   string
	Waktu      time.Time
	SKU        string
	NamaBarang string
	Jumlah     int
	HargaJual  float64
	HargaBeli  float64
	Total      float64
	Laba       float64
}

func GenerateLaporanPenjualan(db *gorm.DB) {
	var barangKeluar []models.BarangKeluar

	db.Order("waktu").Find(&barangKeluar)

}

func generateDetaila(barang models.Barang, barang2Masuk []models.BarangMasuk) NilaiBarangDetail {
	var nilai NilaiBarangDetail
	var totalNilai float64
	var totalTerima int
	for _, barangMasuk := range barang2Masuk {
		totalNilai += barangMasuk.Harga * float64(barangMasuk.JumlahDiterima)
		totalTerima += barangMasuk.JumlahDiterima
	}
	nilai.NamaBarang = barang.Nama
	nilai.Total = totalNilai
	nilai.Jumlah = barang.Jumlah
	nilai.Rata2Harga = totalNilai / float64(totalTerima)

	return nilai
}
