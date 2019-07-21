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
	SKU        string
	NamaBarang string
	Jumlah     int
	Rata2Harga float64
	Total      float64
}

//HargaTotal untuk menyimpan nilai harga dan semua yang berhubungan
type HargaDanTotal struct {
	SKU         string
	Rata2Harga  float64
	TotalNilai  float64
	TotalTerima int
}

func GenerateLaporanNilai(db *gorm.DB) LaporanNilaiBarang {
	var barang2 []models.Barang
	var laporan LaporanNilaiBarang
	mapHarga := GenerateHargaDanTotalNilaiBarang(db)

	db.Where("jumlah > 0").Find(&barang2)

	laporan.TanggalCetak = time.Now()
	for _, barang := range barang2 {
		laporan.JumlahSKU++
		laporan.TotalBarang += barang.Jumlah

		var barang2Masuk []models.BarangMasuk
		db.Where("barang_id = ?", barang.ID).Find(&barang2Masuk)

		nilai := generateDetail(barang, mapHarga[barang.ID])
		laporan.TotalNilai += nilai.Total
		laporan.Detail = append(laporan.Detail, nilai)

	}

	return laporan

}

func generateDetail(barang models.Barang, harga HargaDanTotal) NilaiBarangDetail {
	var nilai NilaiBarangDetail

	nilai.SKU = barang.SKU
	nilai.NamaBarang = barang.Nama
	nilai.Total = harga.TotalNilai
	nilai.Jumlah = harga.TotalTerima
	nilai.Rata2Harga = harga.Rata2Harga

	return nilai
}

//Fungsi ini untuk membuat map yang berisi harga rata2, total nilai, dan total diterima
func GenerateHargaDanTotalNilaiBarang(db *gorm.DB) map[int]HargaDanTotal {
	var barang2 []models.Barang
	mapHarga := make(map[int]HargaDanTotal)

	db.Find(&barang2)

	for _, barang := range barang2 {
		var barang2Masuk []models.BarangMasuk
		db.Where("barang_id = ?", barang.ID).Find(&barang2Masuk)
		var hargaTotal HargaDanTotal

		for _, barangMasuk := range barang2Masuk {
			hargaTotal.TotalNilai += barangMasuk.Harga * float64(barangMasuk.JumlahDiterima)
			hargaTotal.TotalTerima += barangMasuk.JumlahDiterima
		}
		hargaTotal.Rata2Harga = hargaTotal.TotalNilai / float64(hargaTotal.TotalTerima)
		mapHarga[barang.ID] = hargaTotal
	}
	return mapHarga
}
