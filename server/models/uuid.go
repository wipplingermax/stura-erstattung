package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type UUIDPkey struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeltedAt  gorm.DeletedAt `gorm:"index"`
}
