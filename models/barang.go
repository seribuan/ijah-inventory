package models

//Item
type Barang struct {
	SKU    string `gorm:"unique"`
	nama   string
	jumlah int
}
