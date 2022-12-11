package models

import loggable "github.com/daqingshu/gorm-loggable"

type Transcation struct {
	UUIDPkey
	loggable.LoggableModel
	RefundID uint
	Status   string `gorm:"not null"`
}
