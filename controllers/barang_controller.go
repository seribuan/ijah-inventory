package controllers

import (
	"encoding/csv"
	"net/http"
	"sorabel/models"
	"strconv"

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

func (b *BarangController) Export(c *gin.Context) {
	var barang2 []models.Barang
	db.Find(&barang2)
	w := csv.NewWriter(c.Writer)
	var csvData [][]string

	for _, barang := range barang2 {
		csvData = append(csvData, []string{barang.SKU, barang.Nama, strconv.Itoa(barang.Jumlah)})
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=barang.csv")
	c.Header("Content-Type", "application/octet-stream")
	w.WriteAll(csvData)
}
