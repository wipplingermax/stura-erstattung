package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	FirstName           string `gorm:"not null" json:"firstname"`
	LastName            string `gorm:"not null" json:"lastname"`
	MatriculationNumber string `gorm:"not null" json:"matriculationnumber"`
	UniID               string `gorm:"not null" json:"uniid"`
	Email               string `gorm:"not null" json:"email"`
	Phone               string `json:"phone"`
	IBAN                string `gorm:"not null" json:"iban"`
	BIC                 string `json:"bic"`
	AccountOwner        string `gorm:"not null" json:"accountowner"`
	valid               bool   `gorm:"not null"`
	// Refund              Refund
}
