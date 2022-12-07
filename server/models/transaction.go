package models

type Transcation struct {
	UUIDPkey
	Status string `gorm:"not null"`
}
