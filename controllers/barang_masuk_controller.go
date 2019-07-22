package controllers

import (
	"fmt"
	"net/http"
	"sorabel/models"
	"sorabel/services/exporter"
	"strconv"
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
	//TODO: Transaction
	db.Create(&barangMasuk)
	var barang models.Barang
	db.First(&barang, barangMasuk.BarangID)
	barang.Jumlah += barangMasuk.JumlahDiterima
	db.Save(&barang)
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

func (b *BarangMasukController) Export(c *gin.Context) {
	var barang2Masuk []models.BarangMasuk
	db.Preload("Barang").Find(&barang2Masuk)

	var csvData [][]string

	for _, barangMasuk := range barang2Masuk {
		harga := fmt.Sprintf("%.2f", barangMasuk.Harga)
		total := fmt.Sprintf("%.2f", (float64(barangMasuk.JumlahPesan) * barangMasuk.Harga))
		csvData = append(csvData, []string{
			barangMasuk.Waktu.String(),
			barangMasuk.Barang.SKU,
			barangMasuk.Barang.Nama,
			strconv.Itoa(barangMasuk.JumlahPesan),
			strconv.Itoa(barangMasuk.JumlahDiterima),
			harga,
			total,
			barangMasuk.NoKwitansi,
		})
	}
	exporter.ExportCSV(c, csvData)
}
