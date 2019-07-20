package routes

import (
	"sorabel/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	barangController := new(controllers.BarangController)

	router.GET("/barang/:id", barangController.Get)
	//router.POST("/barang", /posting)
	//router.PUT("/barang", putting)
}
