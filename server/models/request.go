package models

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	FirstName           string
	LastName            string
	MatriculationNumber string
	UniID               string
	Email               string
	Phone               string
	IBAN                string
	BIC                 string
	AccountOwner        string
}
