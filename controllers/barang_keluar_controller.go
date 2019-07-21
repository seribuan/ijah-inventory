package controllers

import (
	"net/http"
	"sorabel/models"
	"time"

	"github.com/gin-gonic/gin"
)

type BarangKeluarController struct{}

func (b *BarangKeluarController) Get(c *gin.Context) {
	var barangKeluar models.BarangKeluar
	db.First(&barangKeluar, c.Param("id"))
	c.JSON(http.StatusOK, barangKeluar)
}

func (b *BarangKeluarController) GetAll(c *gin.Context) {
	var barang2keluar []models.BarangKeluar
	db.Find(&barang2keluar)
	c.JSON(http.StatusOK, barang2keluar)
}

func (b *BarangKeluarController) Post(c *gin.Context) {
	var barangKeluar models.BarangKeluar

	if err := c.ShouldBind(&barangKeluar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	barangKeluar.Waktu = time.Now()

	db.Create(&barangKeluar)
	//TODO: Transaction
	var barang models.Barang
	db.First(&barang, barangKeluar.BarangID)
	barang.Jumlah -= barangKeluar.JumlahKeluar
	db.Save(&barang)
	c.JSON(http.StatusOK, barangKeluar)
}

func (b *BarangKeluarController) Put(c *gin.Context) {
	var barangKeluar models.BarangKeluar
	db.First(&barangKeluar, c.Param("id"))
	if err := c.ShouldBind(&barangKeluar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&barangKeluar)
	c.JSON(http.StatusOK, barangKeluar)
}
