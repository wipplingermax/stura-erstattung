package models

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	Status string `gorm:"not null"`
	Refund Refund
}
