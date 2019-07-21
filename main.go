package main

import (
	"sorabel/controllers"
	"sorabel/models"
	"sorabel/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	//importer.Import("Toko Ijah.xlsx")
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")

	defer db.Close()
	db.AutoMigrate(&models.Barang{}, &models.BarangMasuk{}, &models.BarangKeluar{})

	controllers.InitDB(db)
	router := gin.Default()
	routes.AddRoutes(router)
	router.Run()

}
