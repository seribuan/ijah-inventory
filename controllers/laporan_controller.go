package controllers

import (
	"fmt"
	"net/http"
	"sorabel/services"
	"sorabel/services/exporter"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LaporanController struct{}

func (l *LaporanController) GetLaporanNilai(c *gin.Context) {
	c.JSON(http.StatusOK, services.GenerateLaporanNilai(db))
}

func (l *LaporanController) GetLaporanNilaiCSV(c *gin.Context) {
	laporanNilai := services.GenerateLaporanNilai(db)
	var csvData [][]string

	for _, detail := range laporanNilai.Detail2 {
		hargaRata2 := fmt.Sprintf("%.2f", detail.Rata2Harga)
		total := fmt.Sprintf("%.2f", detail.Total)
		csvData = append(csvData, []string{
			detail.SKU,
			detail.NamaBarang,
			strconv.Itoa(detail.Jumlah),
			hargaRata2,
			total,
		})
	}

	exporter.ExportCSV(c, csvData)
}

func (l *LaporanController) GetLaporanPenjualan(c *gin.Context) {
	c.JSON(http.StatusOK, services.GenerateLaporanPenjualan(db))
}

func (l *LaporanController) GetLaporanPenjualanCSV(c *gin.Context) {
	laporanPenjualan := services.GenerateLaporanPenjualan(db)

	var csvData [][]string

	for _, detail := range laporanPenjualan.Detail2 {
		csvData = append(csvData, []string{
			detail.IDPesanan,
			detail.Waktu.String(),
			detail.SKU,
			detail.NamaBarang,
			strconv.Itoa(detail.Jumlah),
			fmt.Sprintf("%.2f", detail.HargaJual),
			fmt.Sprintf("%.2f", detail.Total),
			fmt.Sprintf("%.2f", detail.HargaBeli),
			fmt.Sprintf("%.2f", detail.Laba),
		})
	}
	exporter.ExportCSV(c, csvData)
}
