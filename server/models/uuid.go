package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type UUIDPkey struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeltedAt  gorm.DeletedAt `gorm:"index"`
}
