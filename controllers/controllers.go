package controllers

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB(dbToBeSet *gorm.DB) {
	db = dbToBeSet
}
