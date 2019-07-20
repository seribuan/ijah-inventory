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

func (b *BarangWriter) Write(SKU string, namaBarang string, jumlah int) {
	barang := models.Barang{SKU: SKU, Nama: namaBarang, Jumlah: jumlah}
	fmt.Println(barang)
	//b.db.Create(&barang)
}
