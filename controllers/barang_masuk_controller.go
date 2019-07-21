package controllers

import (
	"net/http"
	"sorabel/models"
	"time"

	"github.com/gin-gonic/gin"
)

type BarangMasukController struct{}

func (b *BarangMasukController) Get(c *gin.Context) {
	var barangMasuk models.BarangMasuk
	db.First(&barangMasuk, c.Param("id"))
	c.JSON(http.StatusOK, barangMasuk)
}

func (b *BarangMasukController) GetAll(c *gin.Context) {
	var barang2masuk []models.BarangMasuk
	db.Find(&barang2masuk)
	c.JSON(http.StatusOK, barang2masuk)
}

func (b *BarangMasukController) Post(c *gin.Context) {
	var barangMasuk models.BarangMasuk

	if err := c.ShouldBind(&barangMasuk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	barangMasuk.Waktu = time.Now()

	db.Create(&barangMasuk)
	c.JSON(http.StatusOK, barangMasuk)
}

func (b *BarangMasukController) Put(c *gin.Context) {
	var barangMasuk models.BarangMasuk
	db.First(&barangMasuk, c.Param("id"))
	if err := c.ShouldBind(&barangMasuk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&barangMasuk)
	c.JSON(http.StatusOK, barangMasuk)
}
