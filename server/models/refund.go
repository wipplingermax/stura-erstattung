package models

import "gorm.io/gorm"

type Refund struct {
	gorm.Model
	FirstName           string `gorm:"not null"`
	LastName            string `gorm:"not null"`
	MatriculationNumber string `gorm:"not null"`
	UniID               string `gorm:"not null"`
	Email               string `gorm:"not null"`
	Phone               string
	IBAN                string `gorm:"not null"`
	BIC                 string
	AccountOwner        string  `gorm:"not null"`
	Request             Request `gorm:"not null"`
}
