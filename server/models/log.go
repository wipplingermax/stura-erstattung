package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	ActorType     string
	Actor         string
	Action        string
	ReferenceType string
	Reference     string
	Message       string
}
