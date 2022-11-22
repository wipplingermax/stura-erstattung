package models

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	Status string `gorm:"not null"`
	Refund Refund
}

func GETRequest(r *Request, id string) (err error) {

}
