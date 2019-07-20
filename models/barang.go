package models

//Item
type Barang struct {
	ID     int
	SKU    string `gorm:"unique"`
	Nama   string
	Jumlah int
}
