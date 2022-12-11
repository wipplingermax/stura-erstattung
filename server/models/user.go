package models

import loggable "github.com/daqingshu/gorm-loggable"

type User struct {
	UUIDPkey
	loggable.LoggableModel
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}
