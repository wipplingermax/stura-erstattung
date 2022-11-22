package models

import "gorm.io/gorm"

var DB *gorm.DB

type Request struct {
	gorm.Model
	Status string `gorm:"not null"`
	Refund Refund
}

func GETRequest(r *Request, id string) {

	DB.First(&r, "id = ?", id)
}
