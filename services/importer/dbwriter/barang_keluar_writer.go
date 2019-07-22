package dbwriter

import (
	"fmt"
	"sorabel/models"
	"time"

	"github.com/jinzhu/gorm"
)

//NewBarangWriter is used to creata a new BarangWriter
func NewBarangKeluarWriter(db *gorm.DB) *BarangKeluarWriter {
	BarangKeluarWriter := new(BarangKeluarWriter)
	BarangKeluarWriter.db = db
	return BarangKeluarWriter
}

type BarangKeluarWriter struct {
	db *gorm.DB
}

func (b *BarangKeluarWriter) Write(waktu time.Time, sku string, jumlahKeluar int, hargaJual float64, catatan string) {
	var barang models.Barang
	b.db.Where("sku = ?", sku).First(&barang)
	barangMasuk := models.BarangKeluar{Waktu: waktu, BarangID: barang.ID, JumlahKeluar: jumlahKeluar, Harga: hargaJual, Catatan: catatan}
	b.db.Create(&barangMasuk)
	fmt.Println(barangMasuk)
}
