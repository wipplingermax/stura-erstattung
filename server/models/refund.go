package models

import "gorm.io/gorm"

var DB *gorm.DB

type Refund struct {
	UUIDPkey
	Status       string `gorm:"not null"`
	Transactions []Transcation
}
