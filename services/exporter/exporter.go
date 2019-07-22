package exporter

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
)

func ExportCSV(c *gin.Context, csvData [][]string) {
	w := csv.NewWriter(c.Writer)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=barang.csv")
	c.Header("Content-Type", "application/octet-stream")
	w.WriteAll(csvData)
}
