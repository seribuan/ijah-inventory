package controllers

import (
	"net/http"
	"sorabel/models"

	"github.com/gin-gonic/gin"
)

type BarangController struct{}

func (b *BarangController) Get(c *gin.Context) {
	barang := models.Barang{}
	db.First(&barang, c.Param("id"))
	c.JSON(http.StatusOK, barang)
}
