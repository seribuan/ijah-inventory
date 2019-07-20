package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"sorabel/models"
	"sorabel/services/importer"
	"sorabel/services/importer/dbwriter"
)

//"github.com/jinzhu/gorm"
//_ "github.com/jinzhu/gorm/dialects/sqlite"

func main() {
	//importer.Import("Toko Ijah.xlsx")
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	db.AutoMigrate(&models.Barang{})

	barangWriter := dbwriter.NewBarangWriter(db)
	importer.Import("Toko Ijah.xlsx", barangWriter)

	//fmt.Printf("%#v", db)

	//if err {
	//	fmt.Errorf("Something went wrong %v:", err)
	//}

}
