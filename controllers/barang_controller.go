package controllers

import (
	"net/http"
	"sorabel/models"

	"github.com/gin-gonic/gin"
)

type BarangController struct{}

func (b *BarangController) Get(c *gin.Context) {
	var barang models.Barang
	db.First(&barang, c.Param("id"))
	c.JSON(http.StatusOK, barang)
}

func (b *BarangController) GetAll(c *gin.Context) {
	var barang2 []models.Barang
	db.Find(&barang2)
	c.JSON(http.StatusOK, barang2)
}

func (b *BarangController) Post(c *gin.Context) {
	barang := models.Barang{SKU: c.PostForm("sku"), Nama: c.PostForm("nama")}
	db.Create(&barang)
	c.JSON(http.StatusOK, barang)
}

func (b *BarangController) Put(c *gin.Context) {
	var barang models.Barang
	db.First(&barang, c.Param("id"))
	barang.SKU = c.PostForm("sku")
	barang.Nama = c.PostForm("nama")
	db.Save(&barang)
	c.JSON(http.StatusOK, barang)
}
