package controller

import (
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func InitController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}
