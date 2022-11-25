package controller

import (
	"server/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.DB
}
