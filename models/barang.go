package models

//Item
type Barang struct {
	SKU    string `gorm:"unique"`
	Nama   string
	Jumlah int
}
