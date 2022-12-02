package models

import "gorm.io/gorm"

var DB *gorm.DB

type Refund struct {
	gorm.Model
	Status       string `gorm:"not null"`
	Transactions []Transcation
}
