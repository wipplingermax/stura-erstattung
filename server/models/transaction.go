package models

import loggable "github.com/daqingshu/gorm-loggable"

type Transcation struct {
	UUIDPkey
	loggable.LoggableModel
	Status string `gorm:"not null"`
}
