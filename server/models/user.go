package models

type User struct {
	UUIDPkey
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}
