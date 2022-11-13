package models

import "gorm.io/gorm"

type FormRequest struct {
	gorm.Model
	Status  string `gorm:"not null"`
	Error   string
	Request Request //ist das hier eine ID oder tatsächlich ein Objekt?
}
