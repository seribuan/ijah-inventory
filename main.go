package main

import (
	"fmt"
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
	db.AutoMigrate(&models.Barang{})

	barang := models.Barang{}
	db.First(&barang, 1)
	fmt.Println(barang)

	controllers.InitDB(db)
	router := gin.Default()
	routes.AddRoutes(router)
	router.Run()

}
