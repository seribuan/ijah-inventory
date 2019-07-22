package main

import (
	"sorabel/controllers"
	"sorabel/models"
	"sorabel/routes"
	"sorabel/services/importer"
	"sorabel/services/importer/dbwriter"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, _ := gorm.Open("sqlite3", "./gorm.db")

	defer db.Close()
	db.AutoMigrate(&models.Barang{}, &models.BarangMasuk{}, &models.BarangKeluar{})

	controllers.InitDB(db)
	barangKeluarWriter := dbwriter.NewBarangKeluarWriter(db)
	importer.ImportBarangKeluar("Toko Ijah.xlsx", barangKeluarWriter)

	//fmt.Println(services.GenerateLaporanPenjualan(db))

	router := gin.Default()
	routes.AddRoutes(router)
	router.Run()

}
