package dbwriter

import (
	"fmt"
	"sorabel/models"

	"github.com/jinzhu/gorm"
)

//NewBarangWriter is used to creata a new BarangWriter
func NewBarangWriter(db *gorm.DB) *BarangWriter {
	barangWriter := new(BarangWriter)
	barangWriter.db = db
	return barangWriter
}

type BarangWriter struct {
	db *gorm.DB
}

func (b *BarangWriter) Write(sku string, namaBarang string, jumlah int) {
	barang := models.Barang{SKU: sku, Nama: namaBarang, Jumlah: jumlah}
	b.db.Create(&barang)
	fmt.Println(barang)
}
