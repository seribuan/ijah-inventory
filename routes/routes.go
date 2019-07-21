package routes

import (
	"sorabel/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	barangController := new(controllers.BarangController)

	router.GET("/barang/:id", barangController.Get)
	router.GET("/barang", barangController.GetAll)
	router.POST("/barang", barangController.Post)
	router.PUT("/barang/:id", barangController.Put)

	barangMasukController := new(controllers.BarangMasukController)
	router.GET("/barang_masuk/:id", barangMasukController.Get)
	router.GET("/barang_masuk", barangMasukController.GetAll)
	router.POST("/barang_masuk", barangMasukController.Post)
	router.PUT("/barang_masuk/:id", barangMasukController.Put)

	barangKeluarController := new(controllers.BarangKeluarController)
	router.GET("/barang_keluar/:id", barangKeluarController.Get)
	router.GET("/barang_keluar", barangKeluarController.GetAll)
	router.POST("/barang_keluar", barangKeluarController.Post)
	router.PUT("/barang_keluar/:id", barangKeluarController.Put)

}
