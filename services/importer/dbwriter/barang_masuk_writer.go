package dbwriter

import (
	"fmt"
	"sorabel/models"
	"time"

	"github.com/jinzhu/gorm"
)

//NewBarangWriter is used to creata a new BarangWriter
func NewBarangMasukWriter(db *gorm.DB) *BarangMasukWriter {
	barangMasukWriter := new(BarangMasukWriter)
	barangMasukWriter.db = db
	return barangMasukWriter
}

type BarangMasukWriter struct {
	db *gorm.DB
}

func (b *BarangMasukWriter) Write(waktu time.Time, sku string, jumlahPesan int, jumlahDiterima int, hargaBeli float64, noKwitansi string, catatan string) {
	var barang models.Barang
	b.db.Where("sku = ?", sku).First(&barang)
	barangMasuk := models.BarangMasuk{Waktu: waktu, BarangID: barang.ID, JumlahPesan: jumlahPesan, JumlahDiterima: jumlahDiterima, Harga: hargaBeli, NoKwitansi: noKwitansi, Catatan: catatan}
	b.db.Create(&barangMasuk)
	fmt.Println(barangMasuk)
}
