package models

import "gorm.io/gorm"

type Transcation struct {
	gorm.Model
	Status string `gorm:"not null"`
}
