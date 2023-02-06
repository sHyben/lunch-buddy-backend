package models

import (
	"github.com/google/uuid"
	"time"
)

// Model is the base model for all models
// It contains the id, created at and updated at timestamps
// It is embedded in all models
// It is not intended to be used directly
// It is not intended to be called by the user
type Model struct {
	ID        uuid.UUID `gorm:"column:id;primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;" json:"updated_at"`
}
