package models

import loggable "github.com/daqingshu/gorm-loggable"

type User struct {
	UUIDPkey
	loggable.LoggableModel
	Role     string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Password string
}
