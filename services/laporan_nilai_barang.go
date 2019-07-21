package services

import (
	"sorabel/models"
	"time"

	"github.com/jinzhu/gorm"
)

type LaporanNilaiBarang struct {
	TanggalCetak time.Time
	JumlahSKU    int
	TotalBarang  int
	TotalNilai   float64
	Detail       []NilaiBarangDetail
}

type NilaiBarangDetail struct {
	NamaBarang string
	Jumlah     int
	Rata2Harga float64
	Total      float64
}

func GenerateLaporanNilai(db *gorm.DB) LaporanNilaiBarang {
	var barang2 []models.Barang
	var laporan LaporanNilaiBarang

	db.Find(&barang2)

	laporan.TanggalCetak = time.Now()
	for _, barang := range barang2 {
		if barang.Jumlah < 1 {
			continue
		}
		laporan.JumlahSKU++
		laporan.TotalBarang += barang.Jumlah

		var barang2Masuk []models.BarangMasuk
		db.Where("barang_id = ?", barang.ID).Find(&barang2Masuk)
		nilai := generateDetail(barang, barang2Masuk)
		laporan.Detail = append(laporan.Detail, nilai)

	}

	return laporan

}

func generateDetail(barang models.Barang, barang2Masuk []models.BarangMasuk) NilaiBarangDetail {
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
