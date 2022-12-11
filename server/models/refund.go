package models

import (
	loggable "github.com/daqingshu/gorm-loggable"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Refund struct {
	UUIDPkey
	loggable.LoggableModel
	Status       string `gorm:"not null"`
	Transactions []Transcation
}
