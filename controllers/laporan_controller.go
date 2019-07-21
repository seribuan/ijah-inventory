package controllers

import (
	"net/http"
	"sorabel/services"

	"github.com/gin-gonic/gin"
)

type LaporanController struct{}

func (l *LaporanController) GetLaporanNilai(c *gin.Context) {
	c.JSON(http.StatusOK, services.GenerateLaporanNilai(db))
}
